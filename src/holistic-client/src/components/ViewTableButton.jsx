import React from 'react';
import AppContext from '../AppContext';

class ViewTableButton extends React.Component { 
    
  render() {
   return (
      <button type="button" className="btn btn-light" onClick={() => this.context.controllers.viewPage(this.context, "TablePage", {"tableName": this.props.text})}>{this.props.text}</button>);
   }
}

ViewTableButton.contextType = AppContext;

export default ViewTableButton;
