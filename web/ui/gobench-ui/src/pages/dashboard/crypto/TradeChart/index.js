import React from 'react'
import PropTypes from 'prop-types'
import { format } from 'd3-format'
import { timeFormat } from 'd3-time-format'
import { ChartCanvas, Chart } from 'react-stockcharts'
import { CandlestickSeries, LineSeries } from 'react-stockcharts/lib/series'
import { XAxis, YAxis } from 'react-stockcharts/lib/axes'
import {
  CrossHairCursor,
  EdgeIndicator,
  CurrentCoordinate,
  MouseCoordinateX,
  MouseCoordinateY
} from 'react-stockcharts/lib/coordinates'
import { LabelAnnotation, Label, Annotate } from 'react-stockcharts/lib/annotation'
import { discontinuousTimeScaleProvider } from 'react-stockcharts/lib/scale'
import { OHLCTooltip, MovingAverageTooltip } from 'react-stockcharts/lib/tooltip'
import { ema } from 'react-stockcharts/lib/indicator'
import { fitWidth } from 'react-stockcharts/lib/helper'
import { last } from 'react-stockcharts/lib/utils'

class CandleStickChartWithAnnotation extends React.Component {
  render () {
    const annotationProps = {
      y: ({ yScale }) => yScale.range()[0],
      tooltip: d => timeFormat('%B')(d.date)
    }

    const margin = { left: 0, right: 50, top: 10, bottom: 20 }
    const height = 400
    const { type, zoomEvent, data: initialData, width, ratio } = this.props
    const gridHeight = height - margin.top - margin.bottom
    const gridWidth = width - margin.left - margin.right

    const showGrid = true
    const xGrid = showGrid ? { innerTickSize: -1 * gridHeight, tickStrokeOpacity: 0.1 } : {}
    const yGrid = showGrid ? { innerTickSize: -1 * gridWidth, tickStrokeOpacity: 0.1 } : {}

    const ema20 = ema()
      .id(0)
      .options({ windowSize: 20 })
      .merge((d, c) => {
        const b = d
        b.ema20 = c
        return b
      })
      .accessor(d => d.ema20)

    const ema50 = ema()
      .id(2)
      .options({ windowSize: 50 })
      .merge((d, c) => {
        const b = d
        b.ema50 = c
        return b
      })
      .accessor(d => d.ema50)

    const calculatedData = ema20(ema50(initialData))
    const xScaleProvider = discontinuousTimeScaleProvider.inputDateAccessor(d => d.date)
    const { data, xScale, xAccessor, displayXAccessor } = xScaleProvider(calculatedData)

    const start = xAccessor(last(data))
    const end = xAccessor(data[Math.max(0, data.length - 150)])
    const xExtents = [start, end]

    return (
      <ChartCanvas
        height={height}
        ratio={ratio}
        width={width}
        margin={margin}
        type={type}
        seriesName='MSFT'
        data={data}
        xScale={xScale}
        xAccessor={xAccessor}
        displayXAccessor={displayXAccessor}
        xExtents={xExtents}
        zoomEvent={zoomEvent}
      >
        <Label
          x={38}
          y={30}
          fontSize={18}
          text='BTC-USD'
          fontFamily='monospace'
          fontWeight='bold'
        />
        <Label
          x={100}
          y={50}
          fontSize={12}
          text='Gobench'
          fontFamily='monospace'
        />
        <Chart id={1} yExtents={[d => [d.high, d.low]]} padding={{ top: 10, bottom: 20 }}>
          <XAxis
            axisAt='bottom'
            orient='bottom'
            {...xGrid}
            zoomEnabled={zoomEvent}
            fontFamily='monospace'
          />
          <MouseCoordinateX
            fontFamily='monospace'
            at='bottom'
            orient='bottom'
            displayFormat={timeFormat('%Y-%m-%d')}
          />
          <MouseCoordinateY
            fontFamily='monospace'
            at='right'
            orient='right'
            displayFormat={format('.2f')}
          />
          <YAxis
            axisAt='right'
            orient='right'
            ticks={5}
            {...yGrid}
            zoomEnabled={zoomEvent}
            fontFamily='monospace'
          />
          <CandlestickSeries
            fill={d => (d.close > d.open ? '#c23f3f' : '#06a45b')}
            stroke={d => (d.close > d.open ? '#c23f3f' : '#06a45b')}
            opacity={1}
          />

          <LineSeries yAccessor={ema20.accessor()} stroke={ema20.stroke()} />
          <LineSeries yAccessor={ema50.accessor()} stroke={ema50.stroke()} />

          <CurrentCoordinate
            yAccessor={ema20.accessor()}
            fill={ema20.stroke()}
            fontFamily='monospace'
          />
          <CurrentCoordinate
            yAccessor={ema50.accessor()}
            fill={ema50.stroke()}
            fontFamily='monospace'
          />

          <EdgeIndicator
            itemType='last'
            orient='right'
            edgeAt='right'
            fontFamily='monospace'
            yAccessor={d => d.close}
            fill={d => (d.close > d.open ? '#c23f3f' : '#06a45b')}
          />
          <OHLCTooltip origin={[0, 0]} fontFamily='monospace' />
          <MovingAverageTooltip
            origin={[5, 60]}
            fontFamily='monospace'
            options={[
              {
                yAccessor: ema20.accessor(),
                type: ema20.type(),
                stroke: ema20.stroke(),
                windowSize: ema20.options().windowSize
              },
              {
                yAccessor: ema50.accessor(),
                type: ema50.type(),
                stroke: ema50.stroke(),
                windowSize: ema50.options().windowSize
              }
            ]}
          />
          <Annotate
            with={LabelAnnotation}
            when={d => d.date.getDate() === 1 /* some condition */}
            usingProps={annotationProps}
          />
        </Chart>
        <CrossHairCursor strokeDasharray='ShortDash' />
      </ChartCanvas>
    )
  }
}

CandleStickChartWithAnnotation.propTypes = {
  data: PropTypes.array.isRequired,
  width: PropTypes.number.isRequired,
  ratio: PropTypes.number.isRequired,
  type: PropTypes.oneOf(['svg', 'hybrid']).isRequired
}

export default fitWidth(CandleStickChartWithAnnotation)
