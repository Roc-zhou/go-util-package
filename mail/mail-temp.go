package mail

const MailTemp = `<html>

<head>
  <meta http-equiv="content-type" content="text/html; charset=utf-8">

  <title></title>

</head>

<body>
  <div style="display:flex;justify-content:center">
    <div
      style="width:500px;margin:50px 0;background-color: #fff;background-clip: border-box;border: 1px solid rgba(0,0,0,.125);border-radius: .25rem;">
      <div style="padding: .75rem 1.25rem;background-color: #dc3545!important;border-bottom: 1px solid
        rgba(0,0,0,.125);border-radius: calc(.25rem - 1px) calc(.25rem - 1px) 0 0;text-align: center;color: #fff;">
        告警
      </div>
      <div style="padding: .75rem 1.25rem;">
        <h4>错误</h4>
        <p>{{errorTip}}</p>
        <hr>
        <h5>详情</h5>
        <div>
          {{errorinfo}}
        </div>
      </div>
      <div
        style="padding: .75rem 1.25rem;background-color: rgba(0,0,0,.03);border-top: 1px solid rgba(0,0,0,.125);border-radius: calc(.25rem - 1px) calc(.25rem - 1px) 0 0;text-align: center;">
        该邮件不可回复
      </div>
    </div>
  </div>
</body>

</html>`
