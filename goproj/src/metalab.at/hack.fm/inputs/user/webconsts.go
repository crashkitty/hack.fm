package uinput

const html = `
<!DOCTYPE HTML>

<html>
	<head>
		<title>HACK.FM</title>
	</head>

	<body>
		<form method="POST">
			<table>
				<tr>
					<td>Post a link</td>
					<td><input type="text" name="ytlink"/></td>
					<td>
						<input type="radio" name="ytmode" value="1"> Hybrid
						<input type="radio" name="ytmode" value="2"> Revolver
						<input type="radio" name="ytmode" value="3"> Queue
					</td>
					<td><button type="submit" name="mediatype" value="yt">ADD</button></td>
				</tr>
				<tr>
					<td>TTS</td>
					<td><input type="text" name="ttstxt"></td>
					<td><button type="submit" name="mediatype" value="tts">Queue</td>
				</tr>
			</table>
		</form>
	</body>
</html>

`
