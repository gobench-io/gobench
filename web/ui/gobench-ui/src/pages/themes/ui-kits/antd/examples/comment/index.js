/* eslint-disable */
import React from 'react'
import { LikeFilled, LikeOutlined } from '@ant-design/icons'
import { Comment, Tooltip, Avatar } from 'antd'
import moment from 'moment'

class AntdCommentExample extends React.Component {
  state = {
    likes: 0,
    dislikes: 0,
    action: null,
  }

  like = () => {
    this.setState({
      likes: 1,
      dislikes: 0,
      action: 'liked',
    })
  }

  dislike = () => {
    this.setState({
      likes: 0,
      dislikes: 1,
      action: 'disliked',
    })
  }

  render() {
    const { likes, dislikes, action } = this.state

    const actions = [
      <span key="comment-basic-like">
        <Tooltip title="Like">
          {action === 'liked' && <LikeFilled onClick={this.like} />}
          {action !== 'liked' && <LikeOutlined onClick={this.like} />}
        </Tooltip>
        <span style={{ paddingLeft: 8, cursor: 'auto' }}>{likes}</span>
      </span>,
      <span key="comment-basic-dislike">
        <Tooltip title="Dislike">
          {action === 'disliked' && <LikeFilled onClick={this.dislike} />}
          {action !== 'disliked' && <LikeOutlined onClick={this.dislike} />}
        </Tooltip>
        <span style={{ paddingLeft: 8, cursor: 'auto' }}>{dislikes}</span>
      </span>,
      <span key="comment-basic-reply-to">Reply to</span>,
    ]

    return (
      <div>
        <h5 className="mb-3">
          <strong>Basic</strong>
        </h5>
        <div className="mb-5">
          <Comment
            actions={actions}
            author={<a>Han Solo</a>}
            avatar={
              <Avatar
                src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
                alt="Han Solo"
              />
            }
            content={
              <p>
                We supply a series of design principles, practical patterns and high quality design
                resources (Sketch and Axure), to help people create their product prototypes
                beautifully and efficiently.
              </p>
            }
            datetime={
              <Tooltip title={moment().format('YYYY-MM-DD HH:mm:ss')}>
                <span>{moment().fromNow()}</span>
              </Tooltip>
            }
          />
        </div>
        <h5 className="mb-3">
          <strong>Nested</strong>
        </h5>
        <div className="mb-5">
          <Comment
            actions={actions}
            author={<a>Han Solo</a>}
            avatar={
              <Avatar
                src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
                alt="Han Solo"
              />
            }
            content={
              <p>
                We supply a series of design principles, practical patterns and high quality design
                resources (Sketch and Axure), to help people create their product prototypes
                beautifully and efficiently.
              </p>
            }
            datetime={
              <Tooltip title={moment().format('YYYY-MM-DD HH:mm:ss')}>
                <span>{moment().fromNow()}</span>
              </Tooltip>
            }
          >
            <Comment
              actions={actions}
              author={<a>Han Solo</a>}
              avatar={
                <Avatar
                  src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
                  alt="Han Solo"
                />
              }
              content={
                <p>
                  We supply a series of design principles, practical patterns and high quality
                  design resources (Sketch and Axure), to help people create their product
                  prototypes beautifully and efficiently.
                </p>
              }
              datetime={
                <Tooltip title={moment().format('YYYY-MM-DD HH:mm:ss')}>
                  <span>{moment().fromNow()}</span>
                </Tooltip>
              }
            />
          </Comment>
        </div>
      </div>
    )
  }
}

export default AntdCommentExample
