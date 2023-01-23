import { useState } from "react";
import "./Header.css";
import { ThemeButton } from "./theme";

export const DisplayCount = (setTitle, count) => {
  setTitle("total: " + count);
  setTimeout(() => {
    setTitle("Kube-FS");
  }, 1000);
  return "total: " + count;
};
export const Header = ({ date, setDate, count, version, theme, setTheme }) => {
  const [color, setColor] = useState("Magenta");
  const [title, setTitle] = useState("Kube-FS");
  return (
    <header className="header" style={{backgroundColor: `${theme === "light" ? "#fff" :"#333"}`}}>
      <div>
        <h1
          role="title"
          onMouseDown={() => DisplayCount(setTitle, count)}
          onMouseOver={() =>
            setColor("#" + Math.floor(Math.random() * 16777215).toString(16))
          }
          style={{ color: color }}
        >
          {title}
          <div style={{ fontSize: "8pt", color: "grey" }}>{version}</div>
        </h1>
      </div>
      
      <input
        style={{backgroundColor: theme === "light" ? "white" : "#333",color: "magenta"}}
        type="text"
        placeholder="Date yymmdd"
        value={date}
        onChange={(e) => setDate(e.target.value)}
      />
      <ThemeButton theme={theme} setTheme={setTheme} />
    </header>
  );
};
