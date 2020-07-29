import React, { useContext } from 'react';
import { get } from 'lodash';

import { AppContext } from '../../context';

const Scenario = () => {
  const appData = useContext(AppContext);
  return <div>
    <textarea
      disabled
      className="application-scenario"
      value={get(appData, 'scenario', '')}
    />
  </div >;
};
export default Scenario;