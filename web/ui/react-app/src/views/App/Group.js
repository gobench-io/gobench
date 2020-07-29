import React, { useEffect, useState, lazy, Suspense } from 'react';
import { get } from 'lodash';
import GoBenchAPI from '../../api/gobench';

const GraphComponent = lazy(() => import('./Graph'));

const loading = () => <p>Loading group...</p>;

function Group({ group, timestamp }) {
  const [graphs, fetchGraphs] = useState([]);
  const [isCollapse, toggleCollapse] = useState(true);

  useEffect(() => {
    if (group && group.id && !isCollapse) {
      GoBenchAPI.getGraphs(group.id).then(res => {
        return fetchGraphs(res);
      })
    }
  }, [group, isCollapse]);
  return (
    <div className="group">
      <div className="group-header clickable"
        onClick={() => toggleCollapse(!isCollapse)}>
        <h3 title={graphs.id || ''} className="group-title">{get(group, 'name', '')}</h3>
        <span className="collapse-button">
          {isCollapse ? 'Expand' : 'Collapse'}
        </span>
      </div>
      <div className={`group-graphs ${isCollapse ? 'collapse' : ''}`}>
        {
          graphs.length > 0 ?
            graphs.map((graph, index) => {
              return !isCollapse && <Suspense key={graph.id || index} fallback={loading()}>
                <GraphComponent graph={graph} timestamp={timestamp} />
              </Suspense>
            })
            : <p className="text-center">Cannot load graphs.</p>
        }
      </div>
    </div>
  );
}

export default Group;
