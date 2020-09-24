import React from 'react'
import { Helmet } from 'react-helmet'

import AntdButtonExample from './examples/button'
import AntdIconExample from './examples/icon'
import AntdGridExample from './examples/grid'
import AntdLayoutExample from './examples/layout'
import AntdAvatarExample from './examples/avatar'
import AntdBadgeExample from './examples/badge'
import AntdCommentExample from './examples/comment'
import AntdCollapseExample from './examples/collapse'
import AntdCarouselExample from './examples/carousel'
import AntdCardExample from './examples/card'
import AntdCalendarExample from './examples/calendar'
import AntdListExample from './examples/list'
import AntdPopoverExample from './examples/popover'
import AntdTreeExample from './examples/tree'
import AntdTooltipExample from './examples/tooltip'
import AntdTimelineExample from './examples/timeline'
import AntdTagExample from './examples/tag'
import AntdTabsExample from './examples/tabs'
import AntdTableExample from './examples/table'
import AntdAutoCompleteExample from './examples/autocomplete'
import AntdCheckboxExample from './examples/checkbox'
import AntdCascaderExample from './examples/cascader'
import AntdDatePickerExample from './examples/datepicker'
import AntdFormExample from './examples/form'
import AntdInputNumberExample from './examples/inputnumber'
import AntdInputExample from './examples/input'
import AntdMentionsExample from './examples/mentions'
import AntdRateExample from './examples/rate'
import AntdRadioExample from './examples/radio'
import AntdSwitchExample from './examples/switch'
import AntdSliderExample from './examples/slider'
import AntdSelectExample from './examples/select'
import AntdTreeSelectExample from './examples/treeselect'
import AntdTransferExample from './examples/transfer'
import AntdTimePickerExample from './examples/timepicker'
import AntdUploadExample from './examples/upload'
import AntdAlertExample from './examples/alert'
import AntdDrawerExample from './examples/drawer'
import AntdModalExample from './examples/modal'
import AntdMessageExample from './examples/message'
import AntdNotificationExample from './examples/notification'
import AntdProgressExample from './examples/progress'
import AntdPopconfirmExample from './examples/popconfirm'
import AntdSpinExample from './examples/spin'
import AntdSkeletonExample from './examples/skeleton'
import AntdAffixExample from './examples/affix'
import AntdBreadcrumbExample from './examples/breadcrumb'
import AntdDropdownExample from './examples/dropdown'
import AntdMenuExample from './examples/menu'
import AntdPaginationExample from './examples/pagination'
import AntdStepsExample from './examples/steps'
import AntdAnchorExample from './examples/anchor'
import AntdBackTopExample from './examples/backtop'
import AntdDividerExample from './examples/divider'

const examples = [
  {
    name: 'Button',
    description: 'To trigger an operation.',
    link: 'https://ant.design/components/button/',
    component: <AntdButtonExample />,
  },
  {
    name: 'Icon',
    description: 'Semantic vector graphics.',
    link: 'https://ant.design/components/icon/',
    component: <AntdIconExample />,
  },
  {
    name: 'Grid',
    description: '24 Grids System.',
    link: 'https://ant.design/components/grid/',
    component: <AntdGridExample />,
  },
  {
    name: 'Layout',
    description: 'Handling the overall layout of a page.',
    link: 'https://ant.design/components/layout/',
    component: <AntdLayoutExample />,
  },
  {
    name: 'Avatar',
    description:
      'Avatars can be used to represent people or objects. It supports images, Icons, or letters.',
    link: 'https://ant.design/components/avatar/',
    component: <AntdAvatarExample />,
  },
  {
    name: 'Badge',
    description: 'Small numerical value or status descriptor for UI elements.',
    link: 'https://ant.design/components/badge/',
    component: <AntdBadgeExample />,
  },
  {
    name: 'Comment',
    description: 'A comment displays user feedback and discussion to website content.',
    link: 'https://ant.design/components/comment/',
    component: <AntdCommentExample />,
  },
  {
    name: 'Collapse',
    description: 'A content area which can be collapsed and expanded.',
    link: 'https://ant.design/components/collapse/',
    component: <AntdCollapseExample />,
  },
  {
    name: 'Carousel',
    description: 'A carousel component. Scales with its container.',
    link: 'https://ant.design/components/carousel/',
    component: <AntdCarouselExample />,
  },
  {
    name: 'Card',
    description: 'Simple rectangular container.',
    link: 'https://ant.design/components/card/',
    component: <AntdCardExample />,
  },
  {
    name: 'Calendar',
    description: 'Container for displaying data in calendar form.',
    link: 'https://ant.design/components/calendar/',
    component: <AntdCalendarExample />,
  },
  {
    name: 'List',
    description: 'Simple List.',
    link: 'https://ant.design/components/list/',
    component: <AntdListExample />,
  },
  {
    name: 'Popover',
    description: 'The floating card popped by clicking or hovering.',
    link: 'https://ant.design/components/popover/',
    component: <AntdPopoverExample />,
  },
  {
    name: 'Tree',
    description: 'Tree structure',
    link: 'https://ant.design/components/tree/',
    component: <AntdTreeExample />,
  },
  {
    name: 'Tooltip',
    description: 'A simple text popup tip.',
    link: 'https://ant.design/components/tooltip/',
    component: <AntdTooltipExample />,
  },
  {
    name: 'Timeline',
    description: 'Vertical display timeline.',
    link: 'https://ant.design/components/timeline/',
    component: <AntdTimelineExample />,
  },
  {
    name: 'Tag',
    description: 'Tag for categorizing or markup.',
    link: 'https://ant.design/components/tag/',
    component: <AntdTagExample />,
  },
  {
    name: 'Tabs',
    description: 'Tabs make it easy to switch between different views.',
    link: 'https://ant.design/components/tabs/',
    component: <AntdTabsExample />,
  },
  {
    name: 'Table',
    description: 'A table displays rows of data.',
    link: 'https://ant.design/components/table/',
    component: <AntdTableExample />,
  },
  {
    name: 'AutoComplete',
    description: 'Autocomplete function of input field.',
    link: 'https://ant.design/components/auto-complete/',
    component: <AntdAutoCompleteExample />,
  },
  {
    name: 'Checkbox',
    description: 'Checkbox component.',
    link: 'https://ant.design/components/checkbox/',
    component: <AntdCheckboxExample />,
  },
  {
    name: 'Cascader',
    description: 'Cascade selection box.',
    link: 'https://ant.design/components/cascader/',
    component: <AntdCascaderExample />,
  },
  {
    name: 'DatePicker',
    description: 'To select or input a date.',
    link: 'https://ant.design/components/date-picker/',
    component: <AntdDatePickerExample />,
  },
  {
    name: 'Form',
    description:
      'Form is used to collect, validate, and submit the user input, usually contains various form items including checkbox, radio, input, select, and etc.',
    link: 'https://ant.design/components/form/',
    component: <AntdFormExample />,
  },
  {
    name: 'InputNumber',
    description: 'Enter a number within certain range with the mouse or keyboard.',
    link: 'https://ant.design/components/input-number/',
    component: <AntdInputNumberExample />,
  },
  {
    name: 'Input',
    description:
      'A basic widget for getting the user input is a text field. Keyboard and mouse can be used for providing or changing data.',
    link: 'https://ant.design/components/input/',
    component: <AntdInputExample />,
  },
  {
    name: 'Mentions',
    description: 'Mentions component.',
    link: 'https://ant.design/components/mention/',
    component: <AntdMentionsExample />,
  },
  {
    name: 'Rate',
    description: 'Rate component.',
    link: 'https://ant.design/components/rate/',
    component: <AntdRateExample />,
  },
  {
    name: 'Radio',
    description: 'Radio component.',
    link: 'https://ant.design/components/radio/',
    component: <AntdRadioExample />,
  },
  {
    name: 'Switch',
    description: 'Switching Selector.',
    link: 'https://ant.design/components/switch/',
    component: <AntdSwitchExample />,
  },
  {
    name: 'Slider',
    description: 'A Slider component for displaying current value and intervals in range.',
    link: 'https://ant.design/components/slider/',
    component: <AntdSliderExample />,
  },
  {
    name: 'Select',
    description: 'Select component to select value from options.',
    link: 'https://ant.design/components/select/',
    component: <AntdSelectExample />,
  },
  {
    name: 'TreeSelect',
    description: 'Tree selection control.',
    link: 'https://ant.design/components/tree-select/',
    component: <AntdTreeSelectExample />,
  },
  {
    name: 'Transfer',
    description: 'Double column transfer choice box.',
    link: 'https://ant.design/components/transfer/',
    component: <AntdTransferExample />,
  },
  {
    name: 'TimePicker',
    description: 'By clicking the input box, you can select a time from a popup panel.',
    link: 'https://ant.design/components/time-picker/',
    component: <AntdTimePickerExample />,
  },
  {
    name: 'Upload',
    description: 'Upload file by selecting or dragging.',
    link: 'https://ant.design/components/upload/',
    component: <AntdUploadExample />,
  },
  {
    name: 'Alert',
    description: 'Alert component for feedback.',
    link: 'https://ant.design/components/alert/',
    component: <AntdAlertExample />,
  },
  {
    name: 'Drawer',
    description: 'Panel slides from screen edge.',
    link: 'https://ant.design/components/drawer/',
    component: <AntdDrawerExample />,
  },
  {
    name: 'Modal',
    description: 'Modal dialogs.',
    link: 'https://ant.design/components/modal/',
    component: <AntdModalExample />,
  },
  {
    name: 'Message',
    description: 'Display global messages as feedback in response to user operations.',
    link: 'https://ant.design/components/message/',
    component: <AntdMessageExample />,
  },
  {
    name: 'Notification',
    description: 'Display a notification message globally.',
    link: 'https://ant.design/components/notification/',
    component: <AntdNotificationExample />,
  },
  {
    name: 'Progress',
    description: 'Display the current progress of an operation flow.',
    link: 'https://ant.design/components/progress/',
    component: <AntdProgressExample />,
  },
  {
    name: 'Popconfirm',
    description: 'A simple and compact confirmation dialog of an action.',
    link: 'https://ant.design/components/popconfirm/',
    component: <AntdPopconfirmExample />,
  },
  {
    name: 'Spin',
    description: 'A spinner for displaying loading state of a page or a section.',
    link: 'https://ant.design/components/spin/',
    component: <AntdSpinExample />,
  },
  {
    name: 'Skeleton',
    description: 'Provide a placeholder at the place which need waiting for loading.',
    link: 'https://ant.design/components/skeleton/',
    component: <AntdSkeletonExample />,
  },
  {
    name: 'Affix',
    description: 'Make an element stick to viewport.',
    link: 'https://ant.design/components/affix/',
    component: <AntdAffixExample />,
  },
  {
    name: 'Breadcrumb',
    description:
      'A breadcrumb displays the current location within a hierarchy. It allows going back to states higher up in the hierarchy.',
    link: 'https://ant.design/components/breadcrumb/',
    component: <AntdBreadcrumbExample />,
  },
  {
    name: 'Dropdown',
    description: 'A dropdown list.',
    link: 'https://ant.design/components/dropdown/',
    component: <AntdDropdownExample />,
  },
  {
    name: 'Menu',
    description: 'Menu list of Navigation.',
    link: 'https://ant.design/components/menu/',
    component: <AntdMenuExample />,
  },
  {
    name: 'Pagination',
    description:
      'A long list can be divided into several pages by Pagination, and only one page will be loaded at a time.',
    link: 'https://ant.design/components/pagination/',
    component: <AntdPaginationExample />,
  },
  {
    name: 'Steps',
    description: 'Steps is a navigation bar that guides users through the steps of a task.',
    link: 'https://ant.design/components/steps/',
    component: <AntdStepsExample />,
  },
  {
    name: 'Anchor',
    description: 'Hyperlinks to scroll on one page.',
    link: 'https://ant.design/components/anchor/',
    component: <AntdAnchorExample />,
  },
  {
    name: 'BackTop',
    description: 'BackTop makes it easy to go back to the top of the page.',
    link: 'https://ant.design/components/back-top/',
    component: <AntdBackTopExample />,
  },
  {
    name: 'Divider',
    description: 'A divider line separates different content.',
    link: 'https://ant.design/components/divider/',
    component: <AntdDividerExample />,
  },
]

class UIKitAntd extends React.Component {
  state = {
    selectedExampleIndex: 0,
  }

  setExample = selectedExampleIndex => {
    this.setState({
      selectedExampleIndex,
    })
  }

  render() {
    const { selectedExampleIndex } = this.state

    const example = examples[selectedExampleIndex]

    return (
      <div>
        <Helmet title="UI Kit / Ant Design" />
        <div className="mb-4">
          {examples.map((item, index) => (
            <button
              key={item.name}
              type="button"
              className={`btn btn-light mr-2 mb-2 ${
                selectedExampleIndex === index ? 'bg-primary text-white' : 'text-primary'
              }`}
              onClick={() => this.setExample(index)}
            >
              {item.name}
            </button>
          ))}
        </div>
        <div className="card">
          <div className="card-header">
            <h5>
              <strong className="mr-3">{example.name}</strong>
              <a
                href={example.link}
                rel="noopener noreferrer"
                target="_blank"
                className="btn btn-sm btn-light mr-3"
              >
                Component Docs & Examples
                <i className="fe fe-corner-right-up" />
              </a>
              <a
                href="https://ant.design/docs/react/introduce"
                rel="noopener noreferrer"
                target="_blank"
                className="btn btn-sm btn-light mr-3"
              >
                Ant Design Docs
                <i className="fe fe-corner-right-up" />
              </a>
            </h5>
            <p className="mb-0">{example.description}</p>
          </div>
          <div className="card-body">
            <div>{example.component}</div>
          </div>
        </div>
      </div>
    )
  }
}

export default UIKitAntd
