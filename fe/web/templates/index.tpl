<!doctype html>
<html>
<head>
  <title>ABC 123 (Tilt Demo)</title>
  <link href="https://fonts.googleapis.com/css?family=Lora:700|Varela+Round" rel="stylesheet">
</head>
<style>
  /* CSS Variables */
  :root {
    --white:        #FFFFFF;
    --white-var1:   #FFFDF6;
    --white-var2:   #FAF6E9;
    --tan:          #ECE8D9;
    --tan-dark:     #B3A16B;
    --yellow:       #C4A827;
    --blue:         #88ABAD;
    --dark-gray:    #19282B;
    /* --blue-green: #678072; */
  }

  html {
    margin: 0;
  }

  body {
    padding: 2em;
    font-family: 'Varela Round', sans-serif;
    color: var(--dark-gray);
    background-color: var(--white-var1);
  }

  header {
    font-family: 'Lora', serif;
    color: var(--tan-dark);
    margin-top: 0;
    margin-left: auto;
    margin-right: auto;
    margin-bottom: 0;
    max-width: 100em;
    padding-top: 1em;
    padding-left: 2em;
    padding-right: 2em;
    font-size: 2em;
    text-align: center;
  }

  .floatr {
    display: inline-block;
    float: right;
  }

  .floatl {
     display: inline-block;
     float: left;
  }

</style>
<body>
  <main>
    <header>Welcome to TILT ✴︎</header>
    <p class="floatl">
        Your letter is: <strong>{{.Letter}}</strong>
    </p>
    <p class="floatr">
        Your number is: <strong>{{.Number}}</strong>
    </p>
  </main>
</body>
</html>
