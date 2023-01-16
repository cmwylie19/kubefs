import './AppContainer.css';

export default function AppContainer(props) {
  return (
    <div className="app-container" role="app-container">
      <div className="header-container">{props.header}</div>
      <div className="images-container">{props.images}</div>
    </div>
  );
}

