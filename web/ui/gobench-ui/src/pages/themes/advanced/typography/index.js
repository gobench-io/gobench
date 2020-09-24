import React from 'react'
import { Helmet } from 'react-helmet'

const AdvancedTypography = () => {
  return (
    <div>
      <Helmet title="Advanced / Typography" />
      <div className="kit__utils__heading">
        <h5>
          <span className="mr-3">Typography</span>
          <a
            href="https://getbootstrap.com/docs/4.3/layout/grid/"
            target="_blank"
            rel="noopener noreferrer"
            className="btn btn-sm btn-light"
          >
            Official Documentation
            <i className="fe fe-corner-right-up" />
          </a>
        </h5>
      </div>
      <section className="card">
        <div className="card-body">
          <div className="row">
            <div className="col-lg-6">
              <div className="mb-5">
                <h5 className="mb-4">
                  <strong>Headings</strong>
                </h5>
                <p className="text-muted">
                  All HTML headings, <code>&lt;h1&gt;</code> through <code>&lt;h6&gt;</code>, are
                  available. <code>.h1</code> through <code>.h6</code> classes are also available,
                  for when you want to match the font styling of a heading but still want your text
                  to be displayed inline
                </p>
                <h1>
                  h1. Bootstrap heading <span className="badge badge-default">LABEL</span>
                </h1>
                <h2>
                  h2. Bootstrap heading <span className="badge badge-primary">LABEL</span>
                </h2>
                <h3>
                  h3. Bootstrap heading <span className="badge badge-info">LABEL</span>
                </h3>
                <h4>
                  h4. Bootstrap heading <span className="badge badge-danger">LABEL</span>
                </h4>
                <h5>
                  h5. Bootstrap heading <span className="badge badge-success">LABEL</span>
                </h5>
                <h6>
                  h6. Bootstrap heading <span className="badge badge-warning">LABEL</span>
                </h6>
              </div>
            </div>
            <div className="col-lg-6">
              <div className="mb-5">
                <h5 className="mb-4">
                  <strong>Styled Headings</strong>
                </h5>
                <p className="text-muted">
                  Create lighter, secondary text in any heading with a generic
                  <code>&lt;small&gt;</code> tag or the <code>.small</code> class
                </p>
                <h1>
                  <i className="fe fe-home mr-3" aria-hidden="true" />
                  Bootstrap heading
                  <span className="text-muted">Secondary text</span>
                </h1>
                <h2>
                  <i className="fe fe-home mr-3" aria-hidden="true" />
                  Bootstrap heading
                  <span className="text-muted">Secondary text</span>
                </h2>
                <h3>
                  <i className="fe fe-home mr-3" aria-hidden="true" />
                  Bootstrap heading
                  <span className="text-muted">Secondary text</span>
                </h3>
                <h4>
                  <i className="fe fe-home mr-3" aria-hidden="true" />
                  Bootstrap heading
                  <span className="text-muted">Secondary text</span>
                </h4>
                <h5>
                  <i className="fe fe-home mr-3" aria-hidden="true" />
                  Bootstrap heading
                  <span className="text-muted">Secondary text</span>
                </h5>
                <h6>
                  <i className="fe fe-home mr-3" aria-hidden="true" />
                  Bootstrap heading
                  <span className="text-muted">Secondary text</span>
                </h6>
              </div>
            </div>
          </div>
          <div className="row">
            <div className="col-lg-4">
              <div className="mb-5">
                <h5 className="mb-4">
                  <strong>Body Copy</strong>
                </h5>
                <p>
                  Bootstrap&apos;s global default <code>font-size</code> is <strong>1rem</strong>,
                  with a<code>line-height</code> of <strong>1.5</strong>. This is applied to the
                  <code>&lt;body&gt;</code> and all paragraphs. In addition,
                  <code>&lt;p&gt;</code> (paragraphs) receive a bottom margin of half their computed
                  line-height (1rem by default)
                </p>
              </div>
            </div>
            <div className="col-lg-4">
              <div className="mb-5">
                <h5 className="mb-4">
                  <strong>Highlight</strong>
                </h5>
                <p>
                  For highlighting a run of text due to its relevance in another context, use the
                  <code>&lt;mark&gt;</code> tag. Like this:
                </p>
                <p>
                  You can use the mark tag to <mark>highlight</mark> text
                </p>
              </div>
            </div>
            <div className="col-lg-4">
              <div className="mb-5">
                <h5 className="mb-4">
                  <strong>Addresses</strong>
                </h5>
                <p>
                  Present contact information for the nearest ancestor or the entire body of work.
                  Preserve formatting by ending all lines with <code>&lt;br&gt;</code>
                </p>
                <address>
                  <strong>Twitter, Inc.</strong>
                  <br />
                  795 Folsom Ave, Suite 600 <br />
                  San Francisco, CA 94107
                  <br />
                  <abbr title="Phone">P:</abbr> (123) 456-7890
                </address>
                <address>
                  <strong>Full Name</strong>
                  <br />
                  <a href="mailto:#">first.last@example.com</a>
                </address>
              </div>
            </div>
          </div>
          <div className="row">
            <div className="col-lg-4">
              <div className="mb-5">
                <h5 className="mb-4">
                  <strong>Lists - Unordered</strong>
                </h5>
                <p className="text-muted">
                  A list of items in which the order does not explicitly matter
                </p>
                <ul>
                  <li>Lorem ipsum dolor sit amet</li>
                  <li>
                    Nulla volutpat aliquam velit
                    <ul>
                      <li>Phasellus iaculis neque</li>
                      <li>Purus sodales ultricies</li>
                      <li>Vestibulum laoreet porttitor sem</li>
                    </ul>
                  </li>
                  <li>Faucibus porta lacus fringilla vel</li>
                  <li>Aenean sit amet erat nunc</li>
                  <li>Eget porttitor lorem</li>
                </ul>
              </div>
            </div>
            <div className="col-lg-4">
              <div className="mb-5">
                <h5 className="mb-4">
                  <strong>Lists - Ordered</strong>
                </h5>
                <p className="text-muted">
                  A list of items in which the order does explicitly matter
                </p>
                <ol>
                  <li>Lorem ipsum dolor sit amet</li>
                  <li>
                    Nulla volutpat aliquam velit
                    <ol>
                      <li>Phasellus iaculis neque</li>
                      <li>Purus sodales ultricies</li>
                      <li>Vestibulum laoreet porttitor sem</li>
                    </ol>
                  </li>
                  <li>Faucibus porta lacus fringilla vel</li>
                  <li>Aenean sit amet erat nunc</li>
                  <li>Eget porttitor lorem</li>
                </ol>
              </div>
            </div>
            <div className="col-lg-4">
              <div className="mb-5">
                <h5 className="mb-4">
                  <strong>Lists - Unstyled</strong>
                </h5>
                <p className="text-muted">
                  Remove the default default styles and left margin on list items. Apply to each
                  nested lists if you need
                </p>
                <ul className="list-unstyled">
                  <li>Lorem ipsum dolor sit amet</li>
                  <li>
                    Nulla volutpat aliquam velit
                    <ul>
                      <li>Phasellus iaculis neque</li>
                      <li>Purus sodales ultricies</li>
                      <li>Vestibulum laoreet porttitor sem</li>
                    </ul>
                  </li>
                  <li>Faucibus porta lacus fringilla vel</li>
                  <li>Aenean sit amet erat nunc</li>
                </ul>
              </div>
            </div>
          </div>
          <div className="row">
            <div className="col-lg-4">
              <div className="mb-5">
                <h5 className="mb-4">
                  <strong>Lists - Icons</strong>
                </h5>
                <p className="text-muted">A list of terms with icons</p>
                <ul className="list-unstyled">
                  <li>
                    <i className="fe fe-check mr-2" aria-hidden="true" />
                    Lorem ipsum dolor sit amet
                  </li>
                  <li>
                    <i className="fe fe-check mr-2" aria-hidden="true" />
                    Nulla volutpat aliquam velit
                    <ul>
                      <li>Phasellus iaculis neque</li>
                      <li>Purus sodales ultricies</li>
                    </ul>
                  </li>
                  <li>
                    <i className="fe fe-check mr-2" aria-hidden="true" />
                    Faucibus porta lacus fringilla vel
                  </li>
                </ul>
              </div>
            </div>
            <div className="col-lg-4 col-sm-6">
              <div className="mb-5">
                <h5 className="mb-4">
                  <strong>Collapsible Submenu</strong>
                </h5>
                <p className="text-muted">A list of terms with icons</p>
                <ul className="list-unstyled">
                  <li>
                    <i className="fe fe-check mr-2" aria-hidden="true" />
                    Lorem ipsum dolor sit amet
                  </li>
                  <li>
                    <a>
                      <i className="fe fe-check mr-2" aria-hidden="false" />
                      Nulla volutpat aliquam velit
                    </a>
                    <ul className="collapse" id="exampleSubmenu" aria-expanded="false">
                      <li>Phasellus iaculis neque</li>
                      <li>Purus sodales ultricies</li>
                    </ul>
                  </li>
                  <li>
                    <i className="fe fe-check mr-2" aria-hidden="true" />
                    Faucibus porta lacus fringilla vel
                  </li>
                </ul>
              </div>
            </div>
            <div className="col-lg-4">
              <div className="mb-5">
                <h5 className="mb-4">
                  <strong>Lists - Inline</strong>
                </h5>
                <p className="text-muted">
                  Place all list items on a single line with <code>display: inline-block;</code> and
                  some light padding
                </p>
                <ul className="list-inline">
                  <li className="list-inline-item">Lorem ipsum</li>
                  <li className="list-inline-item">Phasellus iaculis</li>
                  <li className="list-inline-item">Nulla volutpat</li>
                </ul>
              </div>
            </div>
          </div>
          <div className="row">
            <div className="col-lg-4">
              <div className="mb-5">
                <h5 className="mb-4">
                  <strong>Description</strong>
                </h5>
                <p className="text-muted">A list of terms with their associated descriptions</p>
                <dl>
                  <dt>Description lists</dt>
                  <dd>A description list is perfect for defining terms</dd>
                  <dt>Euismod</dt>
                  <dd>Vestibulum id ligula porta felis euismod semper eget lacinia odio</dd>
                  <dd>Donec id elit non mi porta gravida at eget metus</dd>
                  <dt>Malesuada porta</dt>
                  <dd>Etiam porta sem malesuada magna mollis euismod</dd>
                </dl>
              </div>
            </div>
            <div className="col-lg-8">
              <div className="mb-5">
                <h5 className="mb-4">
                  <strong>Horizontal Description</strong>
                </h5>
                <p className="text-muted">
                  Use the well as a simple effect on an element to give it an inset effect
                </p>
                <dl className="row">
                  <dt className="col-sm-3">Description lists</dt>
                  <dd className="col-sm-9">A description list is perfect for defining terms.</dd>

                  <dt className="col-sm-3">Euismod</dt>
                  <dd className="col-sm-9">
                    Vestibulum id ligula porta felis euismod semper eget lacinia odio sem nec elit.
                  </dd>
                  <dd className="col-sm-9 offset-sm-3">
                    Donec id elit non mi porta gravida at eget metus.
                  </dd>

                  <dt className="col-sm-3">Malesuada porta</dt>
                  <dd className="col-sm-9">Etiam porta sem malesuada magna mollis euismod.</dd>

                  <dt className="col-sm-3 text-truncate">Truncated term is truncated</dt>
                  <dd className="col-sm-9">
                    Fusce dapibus, tellus ac cursus commodo, tortor mauris condimentum nibh, ut
                    fermentum massa justo sit amet risus.
                  </dd>

                  <dt className="col-sm-3">Malesuada porta</dt>
                  <dd className="col-sm-9">Etiam porta sem malesuada magna mollis euismod.</dd>
                </dl>
              </div>
            </div>
          </div>
          <div className="row">
            <div className="col-lg-12">
              <h5 className="mb-4">
                <strong>Blockquotes</strong>
              </h5>
              <p className="text-muted">
                Add a <code>&lt;footer class=&quot;blockquote-footer&quot;&gt;</code> for
                identifying the source. Wrap the name of the source work in{' '}
                <code>&lt;cite&gt;</code>.
              </p>
              <div className="row">
                <div className="col-lg-10 mb-2">
                  <blockquote className="blockquote">
                    <p className="mb-0">
                      Led cursus ante dapibus diam. Sed nisi. Nulla quis sem at nibh elementum
                      imperdiet. Duis sagittis ipsum. Praesent mauris. at nibh elementum imperdiet.
                      Duis sagittis ipsum. Praesent mauris. Fusce nec tellus sed augue semper porta.
                      Mauris massa. Vestibulum lacinia arcu eget nulla. Class aptent taciti sociosqu
                      ad litora torquent per conubia nostra, per inceptos himenaeos. Curabitur
                      sodales ligula in libero. Sed dignissim lacinia nunc.sed cursus ante dapibus
                      diam. Sed nisi. Nulla quis sem at nibh elementum imperdiet
                    </p>
                    <footer className="blockquote-footer">
                      Someone famous in <cite title="Source Title">Source Title</cite>
                    </footer>
                  </blockquote>
                </div>
              </div>
              <div className="row">
                <div className="col-lg-10 offset-lg-2">
                  <blockquote className="blockquote blockquote-reverse">
                    <p className="mb-0">
                      Led cursus ante dapibus diam. Sed nisi. Nulla quis sem at nibh elementum
                      imperdiet. Duis sagittis ipsum. Praesent mauris. at nibh elementum imperdiet.
                      Duis sagittis ipsum. Praesent mauris. Fusce nec tellus sed augue semper porta.
                      Mauris massa. Vestibulum lacinia arcu eget nulla. Class aptent taciti sociosqu
                      ad litora torquent per conubia nostra, per inceptos himenaeos. Curabitur
                      sodales ligula in libero. Sed dignissim lacinia nunc.sed cursus ante dapibus
                      diam. Sed nisi. Nulla quis sem at nibh elementum imperdiet
                    </p>
                    <footer className="blockquote-footer">
                      Someone famous in <cite title="Source Title">Source Title</cite>
                    </footer>
                  </blockquote>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>
    </div>
  )
}

export default AdvancedTypography
