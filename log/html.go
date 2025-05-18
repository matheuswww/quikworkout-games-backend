package server

func GetHtml(msg string) string {
	return `
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
																		<h1 style="color: #cd4646;font-family:system-ui,sans-serif,'Segoe UI', Tahoma, Geneva, Verdana;">Erro no servidor</h1>
																</td>
														</tr>
														<tr>
															<td style="text-align: center; vertical-align: middle;">
																	<div style="color: #b6b6b6; font-family: Arial, Helvetica, sans-serif; font-size: 1.7rem; font-weight: 600; margin: 0 auto;">
																			<p>`+msg+`</p>
																	</div>
															</td>
													</tr>
														<tr>
																<td>
																		<p style="color: #b6b6b6;font-family:sans-serif,'Segoe UI', Tahoma, Geneva, Verdana;">Esta mensagem indica que um erro foi encontrado</p>
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
}