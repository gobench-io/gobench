import React from 'react';
import { get } from 'lodash';
import packageJson from '../../package.json';

const Title = () => <div className="gobench-title">
  <h1 className="app-title">Gobench</h1>
  <span style={{ marginBottom: '8px' }}>A distributed benchmark tool with Golang</span>
  <span>Version: {get(packageJson, 'version', '')}</span>
</div>;

export default Title;