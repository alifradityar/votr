import React from 'react';
import { Link } from 'react-router';

const HeaderNavView = props => {
  return (
    <ul className="nav navbar-nav pull-xs-right">

      <li className="nav-item">
        <Link to="/" className="nav-link">
          List
        </Link>
      </li>

      <li className="nav-item">
        <Link to="/" className="nav-link">
          Create
        </Link>
      </li>

    </ul>
  );
};

class Header extends React.Component {
  render() {
    return (
      <nav className="navbar navbar-light">
        <div className="container">

          <Link to="/" className="navbar-brand">
            {this.props.appName.toLowerCase()}
          </Link>

          <HeaderNavView/>
        </div>
      </nav>
    );
  }
}

export default Header;
