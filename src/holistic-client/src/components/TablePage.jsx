import React from 'react';
import Context from '../Context';
import styled from "styled-components";

class TablePage extends React.Component { 
   render() {
      return <Title>{JSON.stringify(this.props.params)}hi4 {JSON.stringify(this.context.data)}</Title>;
   }
}

const Title = styled.h1`
  font-size: 1.5em;
  text-align: center;
  background-color: ${props => props.theme.fg || "palevioletred"};;
`;


TablePage.contextType = Context;

export default TablePage;
