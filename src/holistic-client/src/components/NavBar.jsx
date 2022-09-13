import React from 'react';
import ViewTableButton from './ViewTableButton';

export default class Navbar extends React.Component { 
    render() {
    
      return <nav className="navbar navbar-expand-lg navbar-light bg-light">
        <div className="navbar-nav">
          <ViewTableButton id="ViewRespositoryTable" text="Respository" context={this.props.context}></ViewTableButton>
          <ViewTableButton id="ViewDatasourceTable" text="Datasource" context={this.props.context}></ViewTableButton>
          <ViewTableButton id="ViewKeyTable" text="Key" context={this.props.context}></ViewTableButton>
          <ViewTableButton id="ViewUserTable" text="User" context={this.props.context}></ViewTableButton>
        </div>
    </nav>;
   }
}
