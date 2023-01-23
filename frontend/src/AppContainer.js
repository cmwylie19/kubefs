import "./AppContainer.css";

export default function AppContainer(props) {
  return (
    <div className="app-container" role="app-container" style={{backgroundColor: `${props.theme === "light" ? "white" : "#333"}`}}>
      <div className="header-container">{props.header}</div>
      <div className="images-container">{props.images}</div>
    </div>
  );
}
