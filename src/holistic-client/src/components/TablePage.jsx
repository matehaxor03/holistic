import React from 'react';

export default class TablePage extends React.Component { 
   render() {
      return <h1>{JSON.stringify(this.props.params)}hi3</h1>;
   }
}
