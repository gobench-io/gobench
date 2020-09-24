/* eslint-disable */
import React from 'react'
import { Checkbox } from 'antd'

function onChange(checkedValues) {
  console.log('checked = ', checkedValues)
}

const plainOptions = ['Apple', 'Pear', 'Orange']
const options = [
  { label: 'Apple', value: 'Apple' },
  { label: 'Pear', value: 'Pear' },
  { label: 'Orange', value: 'Orange' },
]
const optionsWithDisabled = [
  { label: 'Apple', value: 'Apple' },
  { label: 'Pear', value: 'Pear' },
  { label: 'Orange', value: 'Orange', disabled: false },
]

class AntdCheckboxExample extends React.Component {
  render() {
    return (
      <div>
        <div className="row">
          <div className="col-lg-6">
            <h5 className="mb-3">
              <strong>Checkbox Group</strong>
            </h5>
            <Checkbox.Group options={plainOptions} defaultValue={['Apple']} onChange={onChange} />
            <br />
            <br />
            <Checkbox.Group options={options} defaultValue={['Pear']} onChange={onChange} />
            <br />
            <br />
            <Checkbox.Group
              options={optionsWithDisabled}
              disabled
              defaultValue={['Apple']}
              onChange={onChange}
            />
          </div>
          <div className="col-lg-6">
            <h5 className="mb-3">
              <strong>Basic</strong>
            </h5>
            <Checkbox onChange={onChange}>Checkbox</Checkbox>
          </div>
        </div>
      </div>
    )
  }
}

export default AntdCheckboxExample
