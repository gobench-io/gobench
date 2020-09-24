import React from 'react'
import {
  Carousel,
  CarouselItem,
  CarouselControl,
  CarouselIndicators,
  CarouselCaption,
} from 'reactstrap'

const items = [
  {
    src: 'https://via.placeholder.com/1300x500/161537/161537',
    altText: 'Slide 1',
    caption: 'Slide 1',
  },
  {
    src: 'https://via.placeholder.com/1300x500/161537/161537',
    altText: 'Slide 2',
    caption: 'Slide 2',
  },
  {
    src: 'https://via.placeholder.com/1300x500/161537/161537',
    altText: 'Slide 3',
    caption: 'Slide 3',
  },
]

class BootstrapCarouselExample extends React.Component {
  constructor(props) {
    super(props)
    this.state = { activeIndex: 0 }
    this.next = this.next.bind(this)
    this.previous = this.previous.bind(this)
    this.goToIndex = this.goToIndex.bind(this)
    this.onExiting = this.onExiting.bind(this)
    this.onExited = this.onExited.bind(this)
  }

  onExiting() {
    this.animating = true
  }

  onExited() {
    this.animating = false
  }

  next() {
    const { activeIndex } = this.state
    if (this.animating) return
    const nextIndex = activeIndex === items.length - 1 ? 0 : activeIndex + 1
    this.setState({ activeIndex: nextIndex })
  }

  previous() {
    const { activeIndex } = this.state
    if (this.animating) return
    const nextIndex = activeIndex === 0 ? items.length - 1 : activeIndex - 1
    this.setState({ activeIndex: nextIndex })
  }

  goToIndex(newIndex) {
    if (this.animating) return
    this.setState({ activeIndex: newIndex })
  }

  render() {
    const { activeIndex } = this.state

    const slides = items.map(item => {
      return (
        <CarouselItem onExiting={this.onExiting} onExited={this.onExited} key={item.src}>
          <img src={item.src} alt={item.altText} />
          <CarouselCaption captionText={item.caption} captionHeader={item.caption} />
        </CarouselItem>
      )
    })

    return (
      <div>
        <h5 className="mb-4">
          <strong>Default Carousel</strong>
        </h5>
        <div className="mb-5">
          <Carousel activeIndex={activeIndex} next={this.next} previous={this.previous}>
            {slides}
            <CarouselControl
              direction="prev"
              directionText="Previous"
              onClickHandler={this.previous}
            />
            <CarouselControl direction="next" directionText="Next" onClickHandler={this.next} />
          </Carousel>
        </div>
        <h5 className="mb-4">
          <strong>With Indicators</strong>
        </h5>
        <Carousel activeIndex={activeIndex} next={this.next} previous={this.previous}>
          <CarouselIndicators
            items={items}
            activeIndex={activeIndex}
            onClickHandler={this.goToIndex}
          />
          {slides}
          <CarouselControl
            direction="prev"
            directionText="Previous"
            onClickHandler={this.previous}
          />
          <CarouselControl direction="next" directionText="Next" onClickHandler={this.next} />
        </Carousel>
      </div>
    )
  }
}

export default BootstrapCarouselExample
