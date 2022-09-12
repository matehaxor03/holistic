import React from 'react';
import NavbarButton from './NavBarButton';

export default class Navbar extends React.Component { 
    render() {
      return <nav class="navbar navbar-expand-lg navbar-light bg-light">
        <div class="navbar-nav">
          <NavbarButton text="Repositories"></NavbarButton>
        </div>
    </nav>;
   }
}
