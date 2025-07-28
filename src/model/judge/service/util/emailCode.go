package judge_util_service

func EmailCode(title, msg, code string) []byte {
	htmlMessage := `
		<!DOCTYPE html>
		<html>
		<head>
				<meta charset="UTF-8">
		</head>
		<body style="background:black;">
				<div style="text-align: center; width: 100%;">
						<table width="600" cellpadding="0" cellspacing="0" style="margin: 0 auto; max-width: 100%;">
								<tr>
										<td>
												<table width="100%" cellpadding="0" cellspacing="0">
														<tr>
																<td>
																		<h1 style="color: #b6b6b6;font-family:system-ui,sans-serif,'Segoe UI', Tahoma, Geneva, Verdana;">` + title + `</h1>
																</td>
														</tr>
														<tr>
															<td style="text-align: center; vertical-align: middle;">
																	<div style="color: #b6b6b6; font-family: Arial, Helvetica, sans-serif; font-size: 1.7rem; background-color: #11111196; width: 300px; border-radius: 5px; font-weight: 600; border: 1px solid #ffffff0d; margin: 0 auto;letter-spacing: 12px;">
																			<p>` + code + `</p>
																	</div>
															</td>
													</tr>
														<tr>
																<td>
																		<p style="color: #b6b6b6;font-family:sans-serif,'Segoe UI', Tahoma, Geneva, Verdana;">` + msg + `</p>
																		<p style="color: #b6b6b6;font-family:sans-serif,'Segoe UI', Tahoma, Geneva, Verdana;">não foi você? Ignore este email</p>
																		<p style="color: #b6b6b6;font-family:sans-serif,'Segoe UI', Tahoma, Geneva, Verdana;">Atenciosamente,<br>quikworkout</p>
																</td>
														</tr>
												</table>
										</td>
								</tr>
						</table>
				</div>
		</body>
		</html>
		`
	return []byte(htmlMessage)
}
