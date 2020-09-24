const data = {
  accounts: [],
  activationCodes: [],
  ads: [],
  admins: [],
  authorizeddevices: [],
  contacts: [],
  contracts: [],
  devices: [],
  deviceInfo: [{}],
  cabinets: [],
  fotas: [],
  users: [],
  userblocks: [],
  zones: []
}

const { auto, each } = require('async')
const faker = require('faker')
const mockData = async (vu = 100) => {
  await auto(
    {
      users: async () => {
        for (let i = 0; i < vu; i++) {
          data.users.push({
            id: faker.random.uuid(),
            username: faker.internet.userName(),
            password: faker.internet.password(),
            email: faker.internet.email(),
            suspend: faker.random.boolean(),
            admin: faker.random.boolean(),
            emailVerified: faker.random.boolean(),
            verificationToken: faker.internet.password(),
            avatarUrl: faker.internet.avatar(),
            name: faker.name.firstName(),
            birthday: faker.date.past(),
            gender: faker.random.boolean(),
            address: faker.address.streetAddress(),
            phone: faker.phone.phoneNumber(),
            phoneVerified: faker.random.boolean(),
            phoneVerificationCode: faker.internet.password(),
            sendPhoneVerificationAt: faker.date.past(),
            created: faker.date.recent(),
            role: faker.random.arrayElement(['admin', 'manager', 'user'])
          })
        }
        return data.users.map(x => x.id)
      },
      authorizeddevices: async () => {
        const limit = faker.random.number({ min: 300, max: 1000 })
        for (let i = 0; i < limit; i++) {
          data.authorizeddevices.push({
            id: faker.random.uuid(),
            created: faker.date.recent(),
            macAddress: faker.internet.mac(),
            fwName: faker.random.word(),
            username: faker.internet.userName(),
            password: faker.internet.password(),
            pubKey: faker.random.uuid(),
            fwVersion: faker.random.number({ min: 0, max: 10 }),
            fwReportVersionAt: faker.date.recent(),
            tx: faker.random.uuid(),
            rx: faker.random.uuid()
          })
        }
        return data.authorizeddevices.map(x => x.id)
      },
      fotas: async () => {
        const limit = faker.random.number({ min: 10, max: 1000 })
        for (let i = 0; i < limit; i++) {
          data.fotas.push({
            id: faker.random.uuid(),
            created: faker.date.recent(),
            fwName: faker.random.word(),
            version: faker.random.number(),
            signature: faker.internet.password(),
            fwUrl: faker.internet.url()
          })
        }
        return data.fotas.map(x => x.id)
      },
      contracts: async () => {
        const limit = faker.random.number({ min: 10, max: 1000 })
        for (let i = 0; i < limit; i++) {
          data.contracts.push({
            id: faker.random.uuid(),
            created: faker.date.recent(),
            contractNumber: faker.random.number(),
            clientId: faker.random.number(),
            start: faker.date.recent(10),
            end: faker.date.recent(20),
            maxSleepBots: faker.random.number(),
            maxMotionBots: faker.random.number()
          })
        }
        return data.contracts.map(x => x.id)
      },
      zones: async () => {
        const limit = faker.random.number({ min: 10, max: 1000 })
        for (let i = 0; i < limit; i++) {
          data.zones.push({
            id: faker.random.uuid(),
            created: faker.date.recent(),
            name: faker.address.city(),
            lat: faker.address.latitude(),
            long: faker.address.longitude(),
            radius: faker.random.number({ min: 0, max: 100 }),
            address: faker.address.streetAddress(),
            county: faker.address.county,
            city: faker.address.city(),
            country: faker.address.country()
          })
        }
        return data.zones.map(x => x.id)
      },
      userblocks: [
        'users',
        async ({ users }) => {
          const limit = faker.random.number({ min: 1, max: 10 })
          for (let i = 0; i < limit; i++) {
            data.userblocks.push({
              id: faker.random.uuid(),
              created: faker.date.recent(),
              ip: faker.internet.ip(),
              point: faker.random.number(),
              expireFromFirstFail: faker.date.recent(),
              lastFail: faker.date.recent(),
              userId: faker.random.arrayElement(users)
            })
          }
          return data.userblocks.map(x => x.id)
        }
      ],
      devices: [
        'users',
        'authorizeddevices',
        async ({ users, authorizeddevices }) => {
          await each(authorizeddevices, async id => {
            await data.devices.push({
              id,
              created: faker.date.recent(),
              name: faker.random.words(),
              macAddress: faker.internet.mac(),
              username: faker.internet.userName(),
              password: faker.internet.password(),
              userId: faker.random.arrayElement(users)
            })
          })
          return data.devices
        }
      ],
      cabinets: [
        'users',
        'zones',
        'devices',
        async ({ users, zones, devices }) => {
          console.log('cabinet', devices.length)
          const limit = faker.random.number({ min: 100, max: 1000 })
          let _devices = [...devices]
          for (let i = 0; i < limit; i++) {
            const isRegistered = faker.random.boolean()
            const cabinet = {
              id: faker.random.uuid(),
              created: faker.date.recent(),
              name: faker.random.words(),
              userId: faker.random.arrayElement(users),
              zoneId: faker.random.arrayElement(zones),
              dtms: isRegistered ? faker.random.arrayElement(_devices.map(x => x.id)) : null,
              deviceId: isRegistered
                ? faker.random.arrayElement(_devices.map(x => x.macAddress))
                : null,
              wattage: faker.random.number({ min: 0, max: 3000 })
            }
            data.cabinets.push(cabinet)
            _devices = _devices.filter(x => x.id !== cabinet.deviceId)
          }
          console.log('cabbbbb', data.cabinets[0])
          return data.cabinets.map(x => x.id)
        }
      ],
      activationCodes: [
        'users',
        'cabinets',
        'contracts',
        async ({ users, cabinets, contracts }) => {
          const limit = faker.random.number({ min: 10, max: 1000 })
          for (let i = 0; i < limit; i++) {
            data.activationCodes.push({
              id: faker.random.uuid(),
              created: faker.date.recent(),
              activationCode: faker.random.number(),
              userId: faker.random.arrayElement(users),
              cabinetId: faker.random.arrayElement(cabinets),
              contractId: faker.random.arrayElement(contracts)
            })
          }
          return data.activationCodes.map(x => x.id)
        }
      ],
      done: [
        'users',
        'authorizeddevices',
        'fotas',
        'contracts',
        'zones',
        'userblocks',
        'devices',
        'cabinets',
        'activationCodes',
        () => {
          console.log('done', data.cabinets[1])
        }
      ]
    },
    (err, result) => {
      console.log('complete')
      if (err) {
        console.log(err)
        return err
      }
      return data
    }
  )
  return data
}
module.exports = mockData
