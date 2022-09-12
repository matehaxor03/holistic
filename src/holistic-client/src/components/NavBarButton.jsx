import React from 'react';

export default class NavbarButton extends React.Component { 
    
  render() {
      return <button type="button" class="btn btn-light">{this.props.text}</button>;
   }
}
