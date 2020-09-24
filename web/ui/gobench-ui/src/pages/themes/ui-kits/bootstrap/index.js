import React from 'react'
import { Helmet } from 'react-helmet'

import BootstrapAlertsExample from './examples/alerts'
import BootstrapBadgesExample from './examples/badges'
import BootstrapFadeExample from './examples/fade'
import BootstrapFormExample from './examples/form'
import BootstrapTabsExample from './examples/tabs'
import BootstrapToastsExample from './examples/toasts'
import BootstrapTooltipsExample from './examples/tooltips'
import BootstrapButtonDropdownExample from './examples/button-dropdown'
import BootstrapBreadcrumbsExample from './examples/breadcrumbs'
import BootstrapButtonGroupExample from './examples/button-group'
import BootstrapButtonsExample from './examples/buttons'
import BootstrapCardExample from './examples/card'
import BootstrapCarouselExample from './examples/carousel'
import BootstrapCollapseExample from './examples/collapse'
import BootstrapDropdownsExample from './examples/dropdowns'
import BootstrapInputGroupExample from './examples/input-group'
import BootstrapJumbotronExample from './examples/jumbotron'
import BootstrapTablesExample from './examples/tables'
import BootstrapSpinnersExample from './examples/spinners'
import BootstrapProgressExample from './examples/progress'
import BootstrapPopoversExample from './examples/popovers'
import BootstrapPaginationExample from './examples/pagination'
import BootstrapNavExample from './examples/nav'
import BootstrapNavbarExample from './examples/navbar'
import BootstrapModalsExample from './examples/modals'
import BootstrapMediaObjectExample from './examples/media-object'
import BootstrapListGroupExample from './examples/listgroup'
import BootstrapLayoutExample from './examples/layout'

const examples = [
  {
    name: 'Alerts',
    description:
      'Provide contextual feedback messages for typical user actions with the handful of available and flexible alert messages.',
    link: 'https://reactstrap.github.io/components/alerts/',
    component: <BootstrapAlertsExample />,
  },
  {
    name: 'Badges / Pills',
    description: 'Documentation and examples for badges, our small count and labeling component.',
    link: 'https://reactstrap.github.io/components/badge/',
    component: <BootstrapBadgesExample />,
  },
  {
    name: 'Breadcrumbs',
    description:
      'Indicate the current page’s location within a navigational hierarchy that automatically adds separators via CSS.',
    link: 'https://reactstrap.github.io/components/breadcrumbs/',
    component: <BootstrapBreadcrumbsExample />,
  },
  {
    name: 'Button Dropdown',
    description:
      'Toggle contextual overlays for displaying lists of links and more with the Bootstrap dropdown plugin.',
    link: 'https://reactstrap.github.io/components/button-dropdown/',
    component: <BootstrapButtonDropdownExample />,
  },
  {
    name: 'Button Group',
    description:
      'Group a series of buttons together on a single line with the button group, and super-power them with JavaScript.',
    link: 'https://reactstrap.github.io/components/button-group/',
    component: <BootstrapButtonGroupExample />,
  },
  {
    name: 'Buttons',
    description:
      'Use Bootstrap’s custom button styles for actions in forms, dialogs, and more with support for multiple sizes, states, and more.',
    link: 'https://reactstrap.github.io/components/buttons/',
    component: <BootstrapButtonsExample />,
  },
  {
    name: 'Card',
    description:
      'Bootstrap’s cards provide a flexible and extensible content container with multiple variants and options.',
    link: 'https://reactstrap.github.io/components/card/',
    component: <BootstrapCardExample />,
  },
  {
    name: 'Carousel',
    description:
      'A slideshow component for cycling through elements—images or slides of text—like a carousel.',
    link: 'https://reactstrap.github.io/components/carousel/',
    component: <BootstrapCarouselExample />,
  },
  {
    name: 'Collapse',
    description:
      'Toggle the visibility of content across your project with a few classes and our JavaScript plugins.',
    link: 'https://reactstrap.github.io/components/collapse/',
    component: <BootstrapCollapseExample />,
  },
  {
    name: 'Dropdowns',
    description:
      'Toggle contextual overlays for displaying lists of links and more with the Bootstrap dropdown plugin.',
    link: 'https://reactstrap.github.io/components/dropdowns/',
    component: <BootstrapDropdownsExample />,
  },
  {
    name: 'Fade',
    description: 'The content will fade in and out as the button is pressed.',
    link: 'https://reactstrap.github.io/components/fade/',
    component: <BootstrapFadeExample />,
  },
  {
    name: 'Form',
    description:
      'Examples and usage guidelines for form control styles, layout options, and custom components for creating a wide variety of forms.',
    link: 'https://reactstrap.github.io/components/form/',
    component: <BootstrapFormExample />,
  },
  {
    name: 'Input Group',
    description:
      'Easily extend form controls by adding text, buttons, or button groups on either side of textual inputs, custom selects, and custom file inputs.',
    link: 'https://reactstrap.github.io/components/input-group/',
    component: <BootstrapInputGroupExample />,
  },
  {
    name: 'Jumbotron',
    description: 'Lightweight, flexible component for showcasing hero unit style content.',
    link: 'https://reactstrap.github.io/components/jumbotron/',
    component: <BootstrapJumbotronExample />,
  },
  {
    name: 'Layout',
    description:
      'Components and options for laying out your Bootstrap project, including wrapping containers, a powerful grid system, a flexible media object, and responsive utility classes.',
    link: 'https://reactstrap.github.io/components/layout/',
    component: <BootstrapLayoutExample />,
  },
  {
    name: 'ListGroup',
    description:
      'List groups are a flexible and powerful component for displaying a series of content. Modify and extend them to support just about any content within.',
    link: 'https://reactstrap.github.io/components/listgroup/',
    component: <BootstrapListGroupExample />,
  },
  {
    name: 'Media object',
    description:
      'Documentation and examples for Bootstrap’s media object to construct highly repetitive components like blog comments, tweets, and the like.',
    link: 'https://reactstrap.github.io/components/media/',
    component: <BootstrapMediaObjectExample />,
  },
  {
    name: 'Modals',
    description:
      'Use Bootstrap’s JavaScript modal plugin to add dialogs to your site for lightboxes, user notifications, or completely custom content.',
    link: 'https://reactstrap.github.io/components/modals/',
    component: <BootstrapModalsExample />,
  },
  {
    name: 'Navbar',
    description:
      'Documentation and examples for Bootstrap’s powerful, responsive navigation header, the navbar. Includes support for branding, navigation, and more, including support for our collapse plugin.',
    link: 'https://reactstrap.github.io/components/navbar/',
    component: <BootstrapNavbarExample />,
  },
  {
    name: 'Navs',
    description:
      'Documentation and examples for how to use Bootstrap’s included navigation components.',
    link: 'https://reactstrap.github.io/components/navs/',
    component: <BootstrapNavExample />,
  },
  {
    name: 'Pagination',
    description:
      'Documentation and examples for showing pagination to indicate a series of related content exists across multiple pages.',
    link: 'https://reactstrap.github.io/components/pagination/',
    component: <BootstrapPaginationExample />,
  },
  {
    name: 'Popovers',
    description:
      'Documentation and examples for adding Bootstrap popovers, like those found in iOS, to any element on your site.',
    link: 'https://reactstrap.github.io/components/popovers/',
    component: <BootstrapPopoversExample />,
  },
  {
    name: 'Progress',
    description:
      'Documentation and examples for using Bootstrap custom progress bars featuring support for stacked bars, animated backgrounds, and text labels.',
    link: 'https://reactstrap.github.io/components/progress/',
    component: <BootstrapProgressExample />,
  },
  {
    name: 'Spinners',
    description:
      'Indicate the loading state of a component or page with Bootstrap spinners, built entirely with HTML, CSS, and no JavaScript.',
    link: 'https://reactstrap.github.io/components/spinners/',
    component: <BootstrapSpinnersExample />,
  },
  {
    name: 'Tables',
    description:
      'Documentation and examples for opt-in styling of tables (given their prevalent use in JavaScript plugins) with Bootstrap.',
    link: 'https://reactstrap.github.io/components/tables/',
    component: <BootstrapTablesExample />,
  },
  {
    name: 'Tabs / Pills',
    description: "Documentation and examples for how to use Bootstrap's included tabs components.",
    link: 'https://reactstrap.github.io/components/tabs/',
    component: <BootstrapTabsExample />,
  },
  {
    name: 'Toasts',
    description:
      'Push notifications to your visitors with a toast, a lightweight and easily customizable alert message.',
    link: 'https://reactstrap.github.io/components/toasts/',
    component: <BootstrapToastsExample />,
  },
  {
    name: 'Tooltips',
    description:
      'Documentation and examples for adding custom Bootstrap tooltips with CSS and JavaScript using CSS3 for animations and data-attributes for local title storage.',
    link: 'https://reactstrap.github.io/components/tooltips/',
    component: <BootstrapTooltipsExample />,
  },
]

class UIKitBootstrap extends React.Component {
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
        <Helmet title="UI Kit / Bootstrap" />
        {/* <div className="kit__utils__heading">
          <h5>
            <span className="mr-3">Bootstrap UI Kit</span>
            <a
              href="https://reactstrap.github.io"
              rel="noopener noreferrer"
              target="_blank"
              className="btn btn-sm btn-light"
            >
              Official Documentation
              <i className="fe fe-corner-right-up" />
            </a>
          </h5>
        </div> */}
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
                href="https://reactstrap.github.io/"
                rel="noopener noreferrer"
                target="_blank"
                className="btn btn-sm btn-light mr-3"
              >
                Bootstrap Docs
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

export default UIKitBootstrap
