import React, { Component } from 'react';

class Navbar extends Component {
    state = { }
    render () {
        return (<nav className='navbar'>
            <div className="container-fluid">
                <a href="#" className="navbar-brand">Navbar</a>
            </div>
        </nav>);
    }
}

export default Navbar;