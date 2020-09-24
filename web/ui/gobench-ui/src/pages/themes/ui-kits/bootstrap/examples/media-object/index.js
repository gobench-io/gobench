import React from 'react'
import { Media } from 'reactstrap'

class BootstrapMediaObjectExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-4">
          <strong>Default Media Object</strong>
        </h5>
        <Media className="mb-5">
          <Media left>
            <Media
              className="mr-4"
              object
              src="https://via.placeholder.com/64x64/f0f2f4/f0f2f4"
              alt="Generic placeholder image"
            />
          </Media>
          <Media body>
            <Media heading>Media heading</Media>
            Cras sit amet nibh libero, in gravida nulla. Nulla vel metus scelerisque ante
            sollicitudin commodo. Cras purus odio, vestibulum in vulputate at, tempus viverra
            turpis. Fusce condimentum nunc ac nisi vulputate fringilla. Donec lacinia congue felis
            in faucibus.
          </Media>
        </Media>
        <div className="mt-3">
          <Media list>
            <Media tag="li" className="mb-5">
              <Media left>
                <Media
                  className="mr-4"
                  object
                  src="https://via.placeholder.com/64x64/f0f2f4/f0f2f4"
                  alt="Generic placeholder image"
                />
              </Media>
              <Media body>
                <Media heading>Media heading</Media>
                Cras sit amet nibh libero, in gravida nulla. Nulla vel metus scelerisque ante
                sollicitudin commodo. Cras purus odio, vestibulum in vulputate at, tempus viverra
                turpis. Fusce condimentum nunc ac nisi vulputate fringilla. Donec lacinia congue
                felis in faucibus.
                <Media className="mb-5">
                  <Media left>
                    <Media
                      className="mr-4"
                      object
                      src="https://via.placeholder.com/64x64/f0f2f4/f0f2f4"
                      alt="Generic placeholder image"
                    />
                  </Media>
                  <Media body>
                    <Media heading>Nested media heading</Media>
                    Cras sit amet nibh libero, in gravida nulla. Nulla vel metus scelerisque ante
                    sollicitudin commodo. Cras purus odio, vestibulum in vulputate at, tempus
                    viverra turpis. Fusce condimentum nunc ac nisi vulputate fringilla. Donec
                    lacinia congue felis in faucibus.
                    <Media>
                      <Media left>
                        <Media
                          className="mr-4"
                          object
                          src="https://via.placeholder.com/64x64/f0f2f4/f0f2f4"
                          alt="Generic placeholder image"
                        />
                      </Media>
                      <Media body>
                        <Media heading>Nested media heading</Media>
                        Cras sit amet nibh libero, in gravida nulla. Nulla vel metus scelerisque
                        ante sollicitudin commodo. Cras purus odio, vestibulum in vulputate at,
                        tempus viverra turpis. Fusce condimentum nunc ac nisi vulputate fringilla.
                        Donec lacinia congue felis in faucibus.
                      </Media>
                    </Media>
                  </Media>
                </Media>
                <Media>
                  <Media left>
                    <Media
                      className="mr-4"
                      object
                      src="https://via.placeholder.com/64x64/f0f2f4/f0f2f4"
                      alt="Generic placeholder image"
                    />
                  </Media>
                  <Media body>
                    <Media heading>Nested media heading</Media>
                    Cras sit amet nibh libero, in gravida nulla. Nulla vel metus scelerisque ante
                    sollicitudin commodo. Cras purus odio, vestibulum in vulputate at, tempus
                    viverra turpis. Fusce condimentum nunc ac nisi vulputate fringilla. Donec
                    lacinia congue felis in faucibus.
                  </Media>
                </Media>
              </Media>
            </Media>
            <Media tag="li">
              <Media body>
                <Media heading>Media heading</Media>
                Cras sit amet nibh libero, in gravida nulla. Nulla vel metus scelerisque ante
                sollicitudin commodo. Cras purus odio, vestibulum in vulputate at, tempus viverra
                turpis. Fusce condimentum nunc ac nisi vulputate fringilla. Donec lacinia congue
                felis in faucibus.
              </Media>
              <Media right>
                <Media
                  className="ml-4"
                  object
                  src="https://via.placeholder.com/64x64/f0f2f4/f0f2f4"
                  alt="Generic placeholder image"
                />
              </Media>
            </Media>
          </Media>
        </div>
      </div>
    )
  }
}

export default BootstrapMediaObjectExample
