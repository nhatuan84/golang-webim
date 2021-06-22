<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.4.1/css/bootstrap.min.css">
  <style type="text/css">
    *,body {
      margin: 0px;
      padding: 0px;
    }

    body {
      margin: 0px;
      font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
      font-size: 14px;
      line-height: 20px;
      background-color: #fff;
    }

    header,
    footer {
      width: 960px;
      margin-left: auto;
      margin-right: auto;
    }

    .logo {
      background-image: url('/static/img/chat.png');
      background-repeat: no-repeat;
      -webkit-background-size: 100px 100px;
      background-size: 100px 100px;
      background-position: center center;
      text-align: center;
      font-size: 42px;
      padding: 250px 0 70px;
      font-weight: normal;
      text-shadow: 0px 1px 2px #ddd;
    }

    header {
      padding: 100px 0;
    }

    footer {
      line-height: 1.8;
      text-align: center;
      padding: 50px 0;
      color: #999;
    }

    .description {
      text-align: center;
      font-size: 16px;
    }

    a {
      color: #444;
      text-decoration: none;
    }

    .backdrop {
      position: absolute;
      width: 100%;
      height: 100%;
      box-shadow: inset 0px 0px 100px #ddd;
      z-index: -1;
      top: 0px;
      left: 0px;
    }
    .center {
      margin: auto;
      width: 50%;
      border: 5px;
      padding: 10px;
    }

  </style>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8"> 
</head>

<body>
  <header>
    <h1 class="logo">Please Login</h1>
    <div class="center">
          {{if .flash.error}}
          <blockquote>{{.flash.error}}</blockquote>
          {{end}}
          <form method="POST">
            <input type="hidden" name="token" value="{{.token}}">
            <div class="form-group">
              <label for="username">Username</label>
              <input type="text" class="form-control" name="username" placeholder="Username">
            </div>
            <div class="form-group">
              <label for="password">Password</label>
              <input type="password" class="form-control" name="password" placeholder="Password">
            </div>
            <button type="submit" class="btn btn-primary">Submit</button>
          </form>
    </div>
  </header>
</body>
</html>
