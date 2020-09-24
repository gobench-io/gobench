import React from 'react'
import { Pagination, PaginationItem, PaginationLink } from 'reactstrap'

class BootstrapPaginationExample extends React.Component {
  render() {
    return (
      <div>
        <h5 className="mb-4">
          <strong>Default Pagination</strong>
        </h5>
        <div className="mb-5">
          <Pagination aria-label="Page navigation example">
            <PaginationItem>
              <PaginationLink first />
            </PaginationItem>
            <PaginationItem>
              <PaginationLink previous />
            </PaginationItem>
            <PaginationItem active>
              <PaginationLink>1</PaginationLink>
            </PaginationItem>
            <PaginationItem disabled>
              <PaginationLink>2</PaginationLink>
            </PaginationItem>
            <PaginationItem>
              <PaginationLink>3</PaginationLink>
            </PaginationItem>
            <PaginationItem>
              <PaginationLink>4</PaginationLink>
            </PaginationItem>
            <PaginationItem>
              <PaginationLink>5</PaginationLink>
            </PaginationItem>
            <PaginationItem>
              <PaginationLink next />
            </PaginationItem>
            <PaginationItem>
              <PaginationLink last />
            </PaginationItem>
          </Pagination>
        </div>
        <h5 className="mb-4">
          <strong>Sizing</strong>
        </h5>
        <div className="mb-5">
          <Pagination size="lg" aria-label="Page navigation example">
            <PaginationItem>
              <PaginationLink first />
            </PaginationItem>
            <PaginationItem>
              <PaginationLink previous />
            </PaginationItem>
            <PaginationItem>
              <PaginationLink>1</PaginationLink>
            </PaginationItem>
            <PaginationItem>
              <PaginationLink>2</PaginationLink>
            </PaginationItem>
            <PaginationItem>
              <PaginationLink>3</PaginationLink>
            </PaginationItem>
            <PaginationItem>
              <PaginationLink next />
            </PaginationItem>
            <PaginationItem>
              <PaginationLink last />
            </PaginationItem>
          </Pagination>
        </div>
      </div>
    )
  }
}

export default BootstrapPaginationExample
