import React from 'react'

export const ThemeButton = ({theme, setTheme}) => <button style={{borderRadius: "14px",fontFamily:"Red Hat Mono", backgroundColor: "#FF33A5", color: "white"}} onClick={()=>theme === "light" ? setTheme("dark") : setTheme("light") }>{theme}</button>