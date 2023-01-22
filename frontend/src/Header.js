import { useState } from "react";
import "./Header.css";

export const DisplayCount = (setTitle, count) => {
  setTitle("total: " + count);
  setTimeout(() => {
    setTitle("Kube-FS");
  }, 1000);
  return "total: " + count;
};
export const Header = ({ date, setDate, count, version }) => {
  const [color, setColor] = useState("Magenta");
  const [title, setTitle] = useState("Kube-FS");
  return (
    <header className="header">
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
        type="text"
        placeholder="Date yymmdd"
        value={date}
        onChange={(e) => setDate(e.target.value)}
      />
    </header>
  );
};
