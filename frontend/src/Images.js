import { useEffect, useState } from "react";
import * as helpers from "./helpers";
import "./Images.css";

function Images({ date }) {
  const [pics, setPics] = useState(null);

  useEffect(() => {
    helpers.FetchPics(setPics);
    const interval = setInterval(() => {
      helpers.FetchPics(setPics);
    }, 10000);
    return () => clearInterval(interval);

  }, [date]);

  return (
    <>
      {helpers.FilterPics(pics, date) && helpers.FilterPics(pics, date).map((pic) =>
        <img
          alt={pic.Name}
          className="image"
          key={pic.Name}
          src={"http://192.168.1.209:30099" + pic.Path.replace("/media", "")}
          onClick={(e) => helpers.DeletePic(e, pic.Path, pics, setPics, helpers.FilterPics, date)}
          onMouseOver={() => helpers.SetActive(pic.Name, pics, setPics)}
          onMouseOut={() => helpers.SetUnActive(pic.Name, pics, setPics)}
        />)}
    </>
  );
}

export default Images;
