import AppContainer from "./AppContainer";
import { Header } from './Header';
import Images from './Images';
import { useState,useEffect } from 'react';
import * as helpers from "./helpers";

export default function Root () {
    const [date, setDate] = useState("");
    const [pics, setPics] = useState(null);

    useEffect(() => {
      helpers.FetchPics(setPics);
      const interval = setInterval(() => {
        helpers.FetchPics(setPics);
      }, 10000);
      return () => clearInterval(interval);
  
    }, [date]);
  return (
    <AppContainer
      header={<Header date={date} setDate={setDate} count={Array.isArray(pics) ? pics.length : 0} />}
      images={<Images date={date} pic={pics} setPics={setPics} />}
    />
  );
}