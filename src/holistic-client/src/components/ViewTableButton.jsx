import React, { Fragment } from 'react';
import AppContext from '../AppContext';

class ViewTableButton extends React.Component { 
    
  render() {
   //console.log(this.context);
   return (
      <button type="button" className="btn btn-light" onClick={() => this.context.dispatchers.viewPage(this.context, {"type": "TablePage", "tableName": this.props.text})}>{this.props.text}</button>
      );
   }
}

ViewTableButton.contextType = AppContext;

export default ViewTableButton;
