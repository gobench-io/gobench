import React from 'react'
import { Helmet } from 'react-helmet'
import { Tabs } from 'antd'
import { Editor } from 'react-draft-wysiwyg'

const { TabPane } = Tabs

const ExtraAppsGithubDiscuss = () => {
  return (
    <div>
      <Helmet title="Github Discuss" />
      <div className="d-flex flex-wrap">
        <div className="mr-auto pr-3 my-2">
          <i className="fe fe-book font-size-21 mr-2" />
          <div className="text-nowrap d-inline-block font-size-18 text-dark">
            <a className="font-size-18 text-blue" href="#" onClick={e => e.preventDefault()}>
              umijs
            </a>{' '}
            /
            <a className="font-size-18 text-blue" href="#" onClick={e => e.preventDefault()}>
              umi
            </a>
          </div>
        </div>
        <div className="d-flex flex-wrap font-size-16">
          <div className="mr-3 my-2 text-nowrap">
            <i className="fe fe-user-check font-size-21 mr-1" />
            Watch
            <strong className="text-dark font-size-18 ml-1">6,870</strong>
          </div>
          <div className="mr-3 my-2 text-nowrap">
            <i className="fe fe-star font-size-21 mr-1" />
            Star
            <strong className="text-dark font-size-18 ml-1">16,356</strong>
          </div>
          <div className="mr-3 my-2 text-nowrap">
            <i className="fe fe-copy font-size-21 mr-1" />
            Fork
            <strong className="text-dark font-size-18 ml-1">569</strong>
          </div>
        </div>
      </div>
      <Tabs className="kit-tabs-bordered mb-2" defaultActiveKey="1">
        <TabPane tab="Code" key="1" />
        <TabPane
          tab={
            <span>
              Issues <strong>(85)</strong>
            </span>
          }
          key="2"
        />
        <TabPane
          tab={
            <span>
              Pull requests <strong>(4)</strong>
            </span>
          }
          key="3"
        />
        <TabPane tab="Security" key="4" />
        <TabPane tab="Insights" key="5" />
      </Tabs>
      <div className="d-flex flex-xs-wrap border-bottom pb-4 mb-4">
        <div className="mr-auto pr-3">
          <div className="text-dark font-size-24 font-weight-bold mb-2">
            [Feature Request] How to enable custom font that comes from svg #2460
          </div>
          <div className="mb-3">
            <span className="mr-3 text-uppercase badge badge-success">Open</span>
            <a className="font-weight-bold" href="#" onClick={e => e.preventDefault()}>
              zxs2162
            </a>
            wrote this issue 12 days ago · 0 comments
          </div>
        </div>
        <a
          className="btn btn-success align-self-start text-nowrap"
          href="#"
          onClick={e => e.preventDefault()}
        >
          New Issue
        </a>
      </div>
      <div className="row">
        <div className="col-lg-9">
          <div className="d-flex align-items-start mb-3">
            <a
              href="#"
              onClick={e => e.preventDefault()}
              className="kit__utils__avatar kit__utils__avatar--size50 mr-3 flex-shrink-0"
            >
              <img src="resources/images/avatars/5.jpg" alt="Mary Stanform" />
            </a>
            <div className="card flex-grow-1">
              <div className="card-header">
                <a className="font-weight-bold" href="#" onClick={e => e.preventDefault()}>
                  zxs2162
                </a>{' '}
                wrote this issue 12 days ago · 0 comments
              </div>
              <div className="card-body">
                <h3>Description</h3>
                <p>
                  Added child elements to the active text editor lose their scrollTop property when
                  they are scrolled out of view.
                </p>
                <p>
                  Hydrogen a popular atom package adds a React Component inline to the text editor
                  to view results from external code execution through jupyter kernels.
                </p>
                <p>
                  This React Component contains a standard div element (not an immediate child but
                  down the sub tree) that if their is enough output can scroll.
                </p>
                <p>
                  If you scroll inside of the element and then scroll down the text editor for it to
                  fall out of view and be temporarily removed from the DOM. If you scroll back up
                  that scroll position inside the element is lost.
                </p>
                <h3>Steps to Reproduce</h3>
                <ol>
                  <li>Have hydrogen active and setup properly</li>
                  <li>
                    Use <code>Run Cell</code> command from hydrogen on the code below with about 50
                    extra new lines
                  </li>
                </ol>
                <pre>
                  <code>for i in range(1000): print(i)</code>
                </pre>
                <ol start="3">
                  <li>Scroll down the editor</li>
                  <li>Scroll back up the editor</li>
                </ol>
                <p>
                  <strong>Expected behavior:</strong> Everything to look the same
                </p>
                <p>
                  <strong>Actual behavior:</strong> Scroll top of the result view has been set to 0
                  upon removal and adding back to the active DOM
                </p>
                <p>
                  <strong>Reproduces how often:</strong> 100% of the time
                </p>
                <h3>Versions</h3>
                <p>
                  You can get this information from copy and pasting the output of
                  <code>atom --version</code> and <code>apm --version</code> from the command line.
                  Also, please include the OS and what version of the OS you&apos;re running.
                </p>
                <p>
                  Atom : 1.37.0
                  <br />
                  Electron: 2.0.18
                  <br />
                  Chrome : 61.0.3163.100
                  <br />
                  Node : 8.9.3
                </p>
                <h3>Additional Information</h3>
                <p>
                  This isn&apos;t platform specific either, and I don&apos;t believe its a hydrogen
                  problem.
                  <br />I believe this is a problem with how etch processes the render of React
                  Components or how it stores html nodes when removing them. But I understand its
                  more complex than that.
                </p>
              </div>
            </div>
          </div>
          <div className="d-flex align-items-start mb-3">
            <a
              href="#"
              onClick={e => e.preventDefault()}
              className="kit__utils__avatar kit__utils__avatar--size50 mr-3 flex-shrink-0"
            >
              <img src="resources/images/avatars/4.jpg" alt="Mary Stanform" />
            </a>
            <div className="card flex-grow-1">
              <Tabs className="kit-tabs-bordered pt-2 px-3" defaultActiveKey="1">
                <TabPane tab="Write" key="1" />
                <TabPane tab="Preview" key="2" />
              </Tabs>
              <div>
                <Editor
                  toolbarClassName="border-0 px-3"
                  editorClassName="px-3"
                  editorStyle={{
                    height: 250,
                    overflow: 'auto',
                  }}
                />
              </div>
              <div className="card-body border-top py-2 px-3">
                <a
                  href="#"
                  onClick={e => e.preventDefault()}
                  className="btn btn-success btn-with-addon text-nowrap ml-3 my-3"
                >
                  <span className="btn-addon">
                    <i className="btn-addon-icon fe fe-plus-circle" />
                  </span>
                  Add Comment
                </a>
              </div>
            </div>
          </div>
        </div>
        <div className="col-lg-3">
          <div className="py-4 border-bottom">
            <div className="font-weight-bold mb-2">Assignees</div>
            <div>No one assigned</div>
          </div>
          <div className="py-4 border-bottom">
            <div className="font-weight-bold mb-2">Labels</div>
            <div>None yet</div>
          </div>
          <div className="py-4 border-bottom">
            <div className="font-weight-bold mb-2">Projects</div>
            <div>None yet</div>
          </div>
          <div className="py-4 border-bottom">
            <div className="font-weight-bold mb-2">Milestone</div>
            <div>No milestone</div>
          </div>
          <div className="py-4 border-bottom">
            <div className="font-weight-bold mb-2">Notifications</div>
            <a href="#" onClick={e => e.preventDefault()} className="btn btn-light text-blue mb-2">
              Subscribe
            </a>
            <div>You’re not receiving notifications from this thread.</div>
          </div>
          <div className="py-4">
            <div className="font-weight-bold mb-2">4 participants</div>
            <div className="kit__utils__avatarGroup mb-3">
              <div className="kit__utils__avatar kit__utils__avatar--size46">
                <img src="resources/images/avatars/1.jpg" alt="User 1" />
              </div>
              <div className="kit__utils__avatar kit__utils__avatar--size46">
                <img src="resources/images/avatars/2.jpg" alt="User 2" />
              </div>
              <div className="kit__utils__avatar kit__utils__avatar--size46">
                <img src="resources/images/avatars/3.jpg" alt="User 3" />
              </div>
              <div className="kit__utils__avatar kit__utils__avatar--size46">
                <img src="resources/images/avatars/4.jpg" alt="User 4" />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default ExtraAppsGithubDiscuss
