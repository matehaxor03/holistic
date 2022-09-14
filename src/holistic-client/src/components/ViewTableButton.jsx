import React from 'react';
import Context from '../Context';

class ViewTableButton extends React.Component { 
    
  render() {
      return <button type="button" className="btn btn-light" onClick={() => this.context.viewPage("TablePage", {"tableName": this.props.text})}>{this.props.text}</button>;
   }
}

ViewTableButton.contextType = Context;

export default ViewTableButton;
