<html>
<head>
<title></title>
</head>
<body>
<form action="/login" method="post">
	用户名:<input type="text" name="username">
	密码:<input type="password" name="password">
	</br>
	喜爱的水果:
	</br>
	<select name="fruit">
	<option value="apple">apple</option>
	<option value="pear">pear</option>
	<option value="banana">banana</option>
	</select>
	</br>
	性别:
	</br>
	<input type="radio" name="gender" value="1">男
	<input type="radio" name="gender" value="2">女
	<input type="submit" value="登录">

</form>

</body>
</html>