import React from 'react'
import { Helmet } from 'react-helmet'
import { Tabs, Input } from 'antd'

const { TabPane } = Tabs

const ExtraAppsDigitaloceanCreate = () => {
  return (
    <div>
      <Helmet title="DigitalOcean Create" />
      <div className="d-flex flex-wrap align-items-center">
        <div className="kit__utils__avatar kit__utils__avatar--size64 flex-shrink-0 mr-5 mb-3">
          <img src="resources/images/avatars/2.jpg" alt="Mary Stanform" />
        </div>
        <div className="mr-auto mb-3">
          <div className="text-dark font-weight-bold font-size-24">
            <span className="mr-3">Mediatec Software</span>
            <span className="align-middle text-primary text-uppercase font-size-12 badge badge-light">
              Default
            </span>
          </div>
          <div>
            Operational / Developer tooling / Update your project information under Settings
          </div>
        </div>
        <a
          className="btn btn-light btn-lg text-blue font-size-14"
          href="#"
          onClick={e => e.preventDefault()}
        >
          Move Resources →
        </a>
      </div>
      <Tabs className="kit-tabs-bordered mb-3" defaultActiveKey="1">
        <TabPane tab="Resources" key="1" />
        <TabPane tab="Activity" key="2" />
        <TabPane tab="Settings" key="3" />
      </Tabs>
      <h6 className="mb-4 text-uppercase">
        <strong>Choose an image</strong>
      </h6>
      <div className="row mb-5">
        <div className="col-md-3 col-sm-6 col-xs-12">
          <div className="card text-center">
            <div className="card-header pt-3 pb-3">
              <div className="text-uppercase text-dark font-weight-bold">Ubuntu</div>
            </div>
            <div className="card-body pt-3 pb-3">
              <div className="text-center text-gray-5">18.04 x86</div>
            </div>
          </div>
        </div>
        <div className="col-md-3 col-sm-6 col-xs-12">
          <div className="card text-center bg-light border-blue">
            <div className="card-header pt-3 pb-3">
              <div className="text-uppercase text-dark font-weight-bold">Freebsd</div>
            </div>
            <div className="card-body pt-3 pb-3">
              <div className="text-center text-gray-5">18.04 x86</div>
            </div>
          </div>
        </div>

        <div className="col-md-3 col-sm-6 col-xs-12">
          <div className="card text-center">
            <div className="card-header pt-3 pb-3">
              <div className="text-uppercase text-dark font-weight-bold">Fedora</div>
            </div>
            <div className="card-body pt-3 pb-3">
              <div className="text-center text-gray-5">18.04 x86</div>
            </div>
          </div>
        </div>

        <div className="col-md-3 col-sm-6 col-xs-12">
          <div className="card text-center">
            <div className="card-header pt-3 pb-3">
              <div className="text-uppercase text-dark font-weight-bold">Debian</div>
            </div>
            <div className="card-body pt-3 pb-3">
              <div className="text-center text-gray-5">18.04 x86</div>
            </div>
          </div>
        </div>
      </div>
      <h6 className="mb-4 text-uppercase">
        <strong>Choose plan</strong>
      </h6>
      <div className="row">
        <div className="col-md-4">
          <div className="card bg-light border-blue">
            <div className="card-body pt-3 pb-3">
              <div className="text-uppercase text-dark font-weight-bold">Starter</div>
              <div className="text-center text-blue">
                <div className="font-weight-bold font-size-24">Standard Plan</div>
                <div>$0.060 /hour</div>
              </div>
            </div>
          </div>
        </div>
        <div className="col-md-4">
          <div className="card">
            <div className="card-body pt-3 pb-3">
              <div className="text-uppercase text-dark font-weight-bold">Performance</div>
              <div className="text-center text-blue">
                <div className="font-weight-bold font-size-24">General Purpose</div>
                <div>$0.060 /hour</div>
              </div>
            </div>
          </div>
        </div>
        <div className="col-md-4">
          <div className="card">
            <div className="card-body pt-3 pb-3">
              <div className="text-uppercase text-dark font-weight-bold">Pro</div>
              <div className="text-center text-blue">
                <div className="font-weight-bold font-size-24">CPU Optimized</div>
                <div>$0.060 /hour</div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <p className="mb-5">
        Each Droplet adds more free data transfer to your account, starting at 1TB/month and scaling
        with Droplet usage and size. Additional outbound data transfer is billed at $.01/GB.
        <a href="#" onClick={e => e.preventDefault()} className="text-blue">
          Read more
        </a>
        .
      </p>
      <h6 className="mb-4 text-uppercase">
        <strong>Add backups</strong>
      </h6>
      <p className="mb-4">
        Automatic system-level backups. Use the backup images to revert the server or create new
        Droplets. Backups cost 20% of the Droplet price.
      </p>
      <a
        className="btn btn-light btn-lg text-primary font-size-14 d-inline-block mb-5"
        href="#"
        onClick={e => e.preventDefault()}
      >
        Enable Backups
      </a>
      <h6 className="mb-4 text-uppercase">
        <strong>Add block storage</strong>
      </h6>
      <p className="mb-4">
        Block storage lets you add independent storage volumes that can be accessed like local disk
        and moved from one Droplet to another within the same region.
      </p>
      <a
        className="btn btn-light btn-lg text-primary font-size-14 d-inline-block mb-5 text-primary"
        href="#"
        onClick={e => e.preventDefault()}
      >
        Add Volume
      </a>
      <h6 className="mb-4 text-uppercase">
        <strong>Choose a hostname</strong>
      </h6>
      <p className="mb-4">
        Block storage lets you add independent storage volumes that can be accessed like local disk
        and moved from one Droplet to another within the same region.
      </p>
      <div className="row">
        <div className="col-md-5">
          <Input className="width-100p mb-5" placeholder="Please choose droplet hostname..." />
        </div>
      </div>
      <a
        href="#"
        onClick={e => e.preventDefault()}
        className="btn btn-lg btn-success width-100p font-size-16"
      >
        Create Droplet
      </a>
    </div>
  )
}

export default ExtraAppsDigitaloceanCreate
