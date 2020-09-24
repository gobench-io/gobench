import React from 'react'
import { Progress } from 'reactstrap'

class BootstrapProgressExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-4">
          <strong>Default Progress</strong>
        </h5>
        <div className="mb-5">
          <div className="text-center">0%</div>
          <Progress className="mb-4" />
          <div className="text-center">25%</div>
          <Progress className="mb-4" value="25" />
          <div className="text-center">50%</div>
          <Progress className="mb-4" value={50} />
          <div className="text-center">75%</div>
          <Progress className="mb-4" value={75} />
          <div className="text-center">100%</div>
          <Progress className="mb-4" value="100" />
          <div className="text-center">Multiple bars</div>
          <Progress className="mb-4" multi>
            <Progress className="mb-4" bar value="15" />
            <Progress className="mb-4" bar color="success" value="30" />
            <Progress className="mb-4" bar color="info" value="25" />
            <Progress className="mb-4" bar color="warning" value="20" />
            <Progress className="mb-4" bar color="danger" value="5" />
          </Progress>
        </div>
        <h5 className="mb-4">
          <strong>Colors Variants</strong>
        </h5>
        <div className="mb-5">
          <Progress className="mb-4" value={2 * 5} />
          <Progress className="mb-4" color="success" value="25" />
          <Progress className="mb-4" color="info" value={50} />
          <Progress className="mb-4" color="warning" value={75} />
          <Progress className="mb-4" color="danger" value="100" />
        </div>
        <h5 className="mb-4">
          <strong>Labels</strong>
        </h5>
        <div className="mb-5">
          <Progress className="mb-4" value="25">
            25%
          </Progress>
          <Progress className="mb-4" value={50}>
            1/2
          </Progress>
          <Progress className="mb-4" value={75}>
            Youre almost there!
          </Progress>
          <Progress className="mb-4" color="success" value="100">
            You did it!
          </Progress>
          <Progress className="mb-4" multi>
            <Progress className="mb-4" bar value="15">
              Meh
            </Progress>
            <Progress className="mb-4" bar color="success" value="30">
              Wow!
            </Progress>
            <Progress className="mb-4" bar color="info" value="25">
              Cool
            </Progress>
            <Progress className="mb-4" bar color="warning" value="20">
              20%
            </Progress>
            <Progress className="mb-4" bar color="danger" value="5">
              !!
            </Progress>
          </Progress>
        </div>
        <h5 className="mb-4">
          <strong>Striped</strong>
        </h5>
        <div className="mb-5">
          <Progress className="mb-4" striped value={2 * 5} />
          <Progress className="mb-4" striped color="success" value="25" />
          <Progress className="mb-4" striped color="info" value={50} />
          <Progress className="mb-4" striped color="warning" value={75} />
          <Progress className="mb-4" striped color="danger" value="100" />
          <Progress className="mb-4" multi>
            <Progress className="mb-4" striped bar value="10" />
            <Progress className="mb-4" striped bar color="success" value="30" />
            <Progress className="mb-4" striped bar color="warning" value="20" />
            <Progress className="mb-4" striped bar color="danger" value="20" />
          </Progress>
        </div>
        <h5 className="mb-4">
          <strong>Multiple bars / Stacked</strong>
        </h5>
        <div className="mb-5">
          <Progress className="mb-4" animated value={2 * 5} />
          <Progress className="mb-4" animated color="success" value="25" />
          <Progress className="mb-4" animated color="info" value={50} />
          <Progress className="mb-4" animated color="warning" value={75} />
          <Progress className="mb-4" animated color="danger" value="100" />
          <Progress className="mb-4" multi>
            <Progress className="mb-4" animated bar value="10" />
            <Progress className="mb-4" animated bar color="success" value="30" />
            <Progress className="mb-4" animated bar color="warning" value="20" />
            <Progress className="mb-4" animated bar color="danger" value="20" />
          </Progress>
        </div>
        <h5 className="mb-4">
          <strong>Multiple bars / Stacked</strong>
        </h5>
        <div className="mb-5">
          <Progress className="mb-4" multi>
            <Progress className="mb-4" bar value="15" />
            <Progress className="mb-4" bar color="success" value="20" />
            <Progress className="mb-4" bar color="info" value="25" />
            <Progress className="mb-4" bar color="warning" value="20" />
            <Progress className="mb-4" bar color="danger" value="15" />
          </Progress>
          <div className="text-center">With Labels</div>
          <Progress className="mb-4" multi>
            <Progress className="mb-4" bar value="15">
              Meh
            </Progress>
            <Progress className="mb-4" bar color="success" value="35">
              Wow!
            </Progress>
            <Progress className="mb-4" bar color="warning" value="25">
              25%
            </Progress>
            <Progress className="mb-4" bar color="danger" value="25">
              LOOK OUT!!
            </Progress>
          </Progress>
          <div className="text-center">Stripes and Animations</div>
          <Progress className="mb-4" multi>
            <Progress className="mb-4" bar striped value="15">
              Stripes
            </Progress>
            <Progress className="mb-4" bar animated color="success" value="30">
              Animated Stripes
            </Progress>
            <Progress className="mb-4" bar color="info" value="25">
              Plain
            </Progress>
          </Progress>
        </div>
      </div>
    )
  }
}

export default BootstrapProgressExample
