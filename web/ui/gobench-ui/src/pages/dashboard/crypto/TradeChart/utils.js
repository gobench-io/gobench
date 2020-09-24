import { tsvParse } from 'd3-dsv'
import { timeParse } from 'd3-time-format'
import Axios from 'axios'

function parseData (parse) {
  return function cb (d) {
    const b = {}
    b.date = parse(d.date)
    b.open = +d.open
    b.high = +d.high
    b.low = +d.low
    b.close = +d.close
    b.volume = +d.volume
    return b
  }
}

const parseDate = timeParse('%Y-%m-%d')

export default function getData () {
  const promiseMSFT = Axios('//rrag.github.io/react-stockcharts/data/MSFT.tsv')
    .then(response => response.text())
    .then(data => tsvParse(data, parseData(parseDate)))
  return promiseMSFT
}
