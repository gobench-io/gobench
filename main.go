package main

import (
	"context"
	"fmt"
	"math"
	"time"

	gobench "github.com/gobench-io/gobench/dis"
	"github.com/gobench-io/gobench/executor/scenario"
	"github.com/verik-systems/scenario/as-vz/common/logger"
	"github.com/verik-systems/scenario/as-vz/common/utils"
	"github.com/verik-systems/scenario/as-vz/config"
	"github.com/verik-systems/scenario/as-vz/device"
	"github.com/verik-systems/scenario/as-vz/history"
	"github.com/verik-systems/scenario/as-vz/user"
)

const maxServiceIds = 20000

var opts *config.Options
var startTime = time.Now()
var serviceIds []string
var activeUsers map[int]*user.UserObject

const DelayEachApi = 10

// export scenarios
func export() scenario.Vus {
	opts = &config.Options{
		Host:                                  "https://load-test-ecs.dev.originwireless.us",
		MQTTPrefix:                            "this_value_will_be_replaced_after_call_home",
		VirtualUser:                           20000, // number of virtual users
		ActiveUser:                            500,   // 5% of 20k virtual users
		Rate:                                  2,     // virtual users initialize per second
		ExtraTime:                             360.0, // extra time to run a scenario in seconds
		TimeToMakeSureDeviceIsOnlineInMinutes: 15,    // 15 mins
		TimeToMakeSureSetupWMSToBeDone:        10,    // 10 mins
		TTLToken:                              24,    // token TTL in hours
		JwtSignKey:                            "b2d3017c68b12f0d703ba1b2ac79e87c5d4c2209fe3084fe4722ca529ba94d10",
		UC: config.UserConfiguration{
			RangeTimeWithIntervalForMonth: 86400,
			RangeTimeWithIntervalForWeek:  86400,
			RangeTimeWithIntervalForDay:   3600,
			RangeTimeWithIntervalForHour:  300,
			GetHistoryIncludeLevels:       "offlineLevels,engineStopLevels",
			LambdaPoisson:                 12 * 60.0 * 60.0, // 1 day. This is the lambda of poission distribution for user behavior. A user only open app 1 time/day
			// GetServiceConfigInterval:      3600.0,  // 1 hour. This is the lambda of poisson distribution to get service config from server.
			// GetServiceParamInterval:       3600.0,  // 1 hour. This is the lambda of poisson distribution to get service param from server.
			// GetDeviceNetworkInterval:      3600.0,  // 1 hour. This is the lambda of poisson distribution to get device network from server.
			// GetCountDeviceInterval:        3600.0,  // 1 hour. This is the lambda of poisson distribution to get count device from server.
			// GetDeviceStatusInterval:       3600.0,  // 1 hour. This is the lambda of poisson distribution to get device status from server.
			// GetNotiLogInterval:            3600.0,  // 1 hour. This is the lambda of poisson distribution to get notification log from server.
			// GetConatinerInfoInterval:    86400.0, // 1 day. This is the lambda of poisson distribution to get container info from server.
			// GetDataInterval:             86400.0, // 1 day. This is the lambda of poisson distribution to get data from server.
			// NotificationSlienceInterval: 86400.0, // 1 day. This is the lambda of poisson distribution to get notification slience from server.
			// RealtimeInterval:   86400.0, // 1 day. This is the lambda of poisson distribution to the realtime heartbeat cycle.
			// RealtimePeriod:     300.0, // 5 minutes. This is the period for running the realtime heartbeat.
			// RealtimeHBInterval: 300.0, // 5 minutes. This is the interval for sending the realtime heartbeat to the device.
		},
		DC: config.DeviceConfiguration{
			RealtimeInterval:     1.0,     // 1 second. This is the interval for response the realtime heartbeat from a device.
			RealtimeLength:       200.0,   // 200*150% = 300 seconds. This is the period for running the realtime heartbeat from a device.
			HistoryInterval:      300.0,   // 300 seconds. This is the interval which the device will send the history data to the server.
			NotificationInterval: 86400.0, // 24 hours. This is the interval which the device will send the notification data to the server.
		},
	}
	// calculate time to run in minutes
	initVUsInSeconds := float64(opts.VirtualUser)/opts.Rate +
		float64(opts.TimeToMakeSureDeviceIsOnlineInMinutes*60) + float64(opts.TimeToMakeSureDeviceIsOnlineInMinutes*60) +
		float64(opts.TimeToMakeSureSetupWMSToBeDone*60) // init time in seconds

	runningTimeInSeconds := opts.UC.LambdaPoisson // running time

	opts.TimeToRunInMinutes = int64(math.Ceil(initVUsInSeconds+runningTimeInSeconds+opts.ExtraTime) / 60) // time to run in minutes

	if opts.TTLToken*60 < int(opts.TimeToRunInMinutes) {
		logger.Log.Errorf("ttl token must be greater than time to run. please update the ttl token to a larger number.TimeToRunInMinutes: %v, TTLToken: %v", opts.TimeToRunInMinutes, opts.TTLToken*60)
		return nil
	}
	// generate token and set it to opts
	opts.UC.Token = utils.GenerateASPublicToken(opts.JwtSignKey, opts.TTLToken)

	serviceIds = history.GetServiceIdsFromS3()[0:maxServiceIds]
	if len(serviceIds) == 0 {
		logger.Log.Errorf("no pre-populated services were found in historical data. please pre-populate services first")
		return nil
	}
	activeUsers = make(map[int]*user.UserObject, opts.ActiveUser)

	logger.Log.Infof("starting run scenario with options: %+v\n", opts)
	logger.Log.Infof("estimation time: init %v users in %v. running time desition: %v. extratime is %v. Total: %v\n", opts.VirtualUser, initVUsInSeconds/60, runningTimeInSeconds/60, opts.ExtraTime/60, opts.TimeToRunInMinutes)

	return scenario.Vus{
		{
			Nu:   opts.VirtualUser,
			Rate: opts.Rate,
			Fu:   f,
		},
	}
}

func f(ctx context.Context, vui int) {
	// Index begin from 0 -> no. user begin from 1
	vu := vui + 1

	// init new user
	u := user.New(opts, vu, fmt.Sprintf("vu::%06d", vu))

	if err := userInitilize(ctx, u); err != nil {
		logger.Log.Errorf("Failed to initialize user %v: %v", u.GetConatinerId(), err)
		return
	}

	// 2. Device Creation
	d := device.New(
		opts,
		vu,
		fmt.Sprintf("moId::%s", u.Store.MoId),
		u.Store.UA.ContainerId,
		u.Store.UA.STO,
		u.Store.ServiceId,
	)

	d.Run(ctx)
	// delay for everything is ready
	gobench.SleepRateLinear(1 / float64(opts.TimeToMakeSureDeviceIsOnlineInMinutes*60)) // sleep for TimeToMakeSureDeviceIsOnline minutes

	if u.Store.MoId == "" || u.Store.ServiceId == "" {
		logger.Log.Errorf("Failed to get MoId or ServiceId for user %v", u.GetConatinerId())
		return
	}

	ok := deviceSetup(ctx, u, d)
	if !ok {
		return
	}
	//logger.Log.Infof("Device setup completed.. wait for sometime to reduce the traffic before we make the User API Calls \n", )
	gobench.SleepRateLinear(1 / float64(opts.TimeToMakeSureDeviceIsOnlineInMinutes*60))
	//logger.Log.Infof("Wait of 5 mins is completed.. can start making User API calls \n", )

	if vu <= opts.ActiveUser {
		activeUsers[vu] = u
	}

	// wait all users and devices are initialized
	if vu == opts.VirtualUser {

		//user behavior
		for _, u := range activeUsers {
			logger.Log.Infof("active user: %+v\n", *u)
			go everythingIsReady(ctx, u)
		}
	}

	// sleep for TimeToRunInMinutes minutes
	nowTime := time.Now()
	minutesHavePassed := int64(math.Ceil(nowTime.Sub(startTime).Minutes()))
	minutesRemainingTimeToRun := opts.TimeToRunInMinutes - minutesHavePassed
	if minutesRemainingTimeToRun > 0 {
		time.Sleep(time.Minute * time.Duration(minutesRemainingTimeToRun))
	}
}

// userInitilize will do following things:
func userInitilize(ctx context.Context, u *user.UserObject) (err error) {
	if err = u.RequestCreateAccount(ctx); err != nil {
		logger.Log.Errorf("Failed to create account for user %v: %v", u.GetConatinerId(), err)
		return
	}

	if err = u.RequestGetAccount(ctx); err != nil {
		logger.Log.Errorf("RequestGetAccount error for user %v: %v", u.GetConatinerId(), err)
		return
	}

	if err = u.RequestCreateService(ctx); err != nil {
		logger.Log.Errorf("RequestCreateService error for user %v: %v", u.GetConatinerId(), err)
		return
	}

	return
}

// deviceSetup will do following things:
func deviceSetup(ctx context.Context, u *user.UserObject, d *device.DeviceObject) (ok bool) {
	ok = true
	if err := u.RequestSetUpWms(ctx); err != nil {
		logger.Log.Errorf("RequestSetUpWms error for user %v: %v", u.GetConatinerId(), err)
		ok = false
		return
	}

	var retriesGetConfig int = 3
	// waiting for wms setup to be done and device attached to service
	for {
		time.Sleep(time.Duration(opts.TimeToMakeSureSetupWMSToBeDone) * time.Minute)
		// get service config
		isRetry := retriesGetConfig != 3
		if err := u.RequestGetConfigs(ctx, isRetry); err != nil {
			retriesGetConfig--
			continue
		}

		if u.Store.DeviceAttachedService {
			break
		}
		retriesGetConfig--
		if retriesGetConfig == 0 {
			logger.Log.Errorf("Failed to get service config for user %v. Exit!", u.GetConatinerId())
			ok = false
			return
		}
	}

	// automatically generate report, historical, notification, etc.
	go d.ReportHistory(ctx)
	go d.ReportRealtime(ctx)
	go d.ReportNotification(ctx)

	return
}

func everythingIsReady(ctx context.Context, u *user.UserObject) {
	for {
		logger.Log.Infof("Inside the function to make the User API's: %+v\n", *u)
		gobench.SleepRatePoisson(1 / opts.UC.LambdaPoisson)
		select {
		case <-ctx.Done():
			logger.Log.Infof("Returning from the function: %+v\n", *u)
			return
		default:
			logger.Log.Infof("Beginning to make User API calls: %+v\n", *u)
			// 1. Count device
			if err := u.RequestGetCountDevice(ctx); err != nil {
				logger.Log.Errorf("RequestGetCountDevice error for user %v: %v", u.GetConatinerId(), err)
				break
			}

			// 2. Post heartbeat
			time.Sleep(DelayEachApi * time.Second)
			if err := u.RequestCreateHeartbeat(ctx); err != nil {
				logger.Log.Errorw("FAILED request create heartbeat", "user", u.GetConatinerId(), "error", err)
				break
			}
			// 3. Post notification silence
			time.Sleep(DelayEachApi * time.Second)
			if err := u.RequestPostSilence(ctx); err != nil {
				logger.Log.Errorf("RequestPostSilence error for user %v: %v", u.GetConatinerId(), err)
				break
			}

			// 4. Get historical data
			time.Sleep(DelayEachApi * time.Second)
			getHistories(ctx)

			// 6. Get service param
			time.Sleep(DelayEachApi * time.Second)
			if err := u.RequestGetParams(ctx); err != nil {
				logger.Log.Errorf("RequestGetParams error for user %v: %v", u.GetConatinerId(), err)
				break
			}

			// 7. Get device network
			time.Sleep(DelayEachApi * time.Second)
			if err := u.RequestGetDeviceNetwork(ctx); err != nil {
				logger.Log.Errorf("RequestGetDeviceNetwork error for user %v: %v", u.GetConatinerId(), err)
				break
			}

			// 8. Get device status
			time.Sleep(DelayEachApi * time.Second)
			if err := u.RequestGetDeviceStatus(ctx); err != nil {
				logger.Log.Errorf("RequestGetDeviceStatus error for user %v: %v", u.GetConatinerId(), err)
				break
			}

			// 9. Get notification log
			time.Sleep(DelayEachApi * time.Second)
			if err := u.RequestGetNotificationLogs(ctx); err != nil {
				logger.Log.Errorf("RequestGetNotificationLogs error for user %v: %v", u.GetConatinerId(), err)
				break
			}

			// 10. Get container
			time.Sleep(DelayEachApi * time.Second)
			if err := u.RequestGetAccount(ctx); err != nil {
				logger.Log.Errorf("RequestGetAccount error for user %v: %v", u.GetConatinerId(), err)
				break
			}
		}
	}
}

func getHistories(ctx context.Context) {
	startDate := time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2022, 6, 30, 0, 0, 0, 0, time.UTC)
	days := int(endDate.Sub(startDate).Hours() / 24)

	randIndexServiceId := utils.RandomMinMax(0, len(serviceIds))
	serviceId := serviceIds[randIndexServiceId]

	if serviceId == "" {
		logger.Log.Warn("serviceId is empty")
		return
	}

	h := history.NewHistory(serviceId, opts.Host, opts.UC.Token)

	// get histories by hourly
	time.Sleep(DelayEachApi * time.Second)
	ran := utils.RandomMinMax(0, 24*days)
	start := startDate.Add(time.Hour * time.Duration(ran))
	end := start.Add(time.Hour)
	if err := h.RequestGetHistories(ctx, "hourly", start.Unix(), end.Unix(), opts.UC.RangeTimeWithIntervalForHour, opts.UC.GetHistoryIncludeLevels, ""); err != nil {
		return
	}

	// get histories by daily
	time.Sleep(DelayEachApi * time.Second)
	ran = utils.RandomMinMax(0, days)
	start = utils.BeginningOfDate(startDate.AddDate(0, 0, ran))
	end = utils.EndingOfDate(start)
	if err := h.RequestGetHistories(ctx, "daily", start.Unix(), end.Unix(), opts.UC.RangeTimeWithIntervalForDay, opts.UC.GetHistoryIncludeLevels, ""); err != nil {
		return
	}

	// get histories by weekly
	time.Sleep(DelayEachApi * time.Second)
	ran = utils.RandomMinMax(0, days-7)
	start = utils.BeginningOfDate(startDate.AddDate(0, 0, ran))
	end = utils.EndingOfDate(start.AddDate(0, 0, 6))
	if err := h.RequestGetHistories(ctx, "weekly", start.Unix(), end.Unix(), opts.UC.RangeTimeWithIntervalForWeek, opts.UC.GetHistoryIncludeLevels, ""); err != nil {
		return
	}

	// get histories by monthly
	time.Sleep(DelayEachApi * time.Second)
	ran = utils.RandomMinMax(0, days-30)
	start = utils.BeginningOfDate(startDate.AddDate(0, 0, ran))
	end = utils.EndingOfDate(start.AddDate(0, 0, 29))
	if err := h.RequestGetHistories(ctx, "monthly", start.Unix(), end.Unix(), opts.UC.RangeTimeWithIntervalForMonth, opts.UC.GetHistoryIncludeLevels, ""); err != nil {
		return
	}

}
