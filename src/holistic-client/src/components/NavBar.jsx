import React from 'react';
import ViewTableButton from './ViewTableButton';

export default class Navbar extends React.Component { 
    render() {
    
      return <nav className="navbar navbar-expand-lg navbar-light bg-light">
        <div className="navbar-nav">
          <ViewTableButton id="ViewRespositoryTable" text="Respository"></ViewTableButton>
          <ViewTableButton id="ViewDatasourceTable" text="Datasource"></ViewTableButton>
          <ViewTableButton id="ViewKeyTable" text="Key"></ViewTableButton>
          <ViewTableButton id="ViewUserTable" text="User"></ViewTableButton>
        </div>
    </nav>;
   }
}
