import React from 'react';
import { createDevTools } from 'redux-devtools';
import SliderMonitor from 'redux-slider-monitor';
import DiffMonitor from 'redux-devtools-diff-monitor';
import LogMonitor from 'redux-devtools-log-monitor';
import DockMonitor from 'redux-devtools-dock-monitor';
import Inspector from 'redux-devtools-inspector';
import ChartMonitor from 'redux-devtools-chart-monitor';
import Dispatcher from 'redux-devtools-dispatch';
import FilterableLogMonitor from 'redux-devtools-filterable-log-monitor'
const DevTools = createDevTools(
    <DockMonitor toggleVisibilityKey='ctrl-h'
                 changePositionKey='ctrl-q'
                 defaultPosition='bottom'
                 changeMonitorKey="ctrl-x"
                >
          <Inspector />
          <ChartMonitor/>
          <SliderMonitor keyboardEnabled />
    </DockMonitor>
  );
  
  export default DevTools;