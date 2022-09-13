import React from 'react';

export default class ViewTableButton extends React.Component { 
    
  render() {
      return <button type="button" className="btn btn-light" onClick={() => this.props.context.viewPage("TablePage", {"tableName": this.props.text})}>{this.props.text}</button>;
   }
}
