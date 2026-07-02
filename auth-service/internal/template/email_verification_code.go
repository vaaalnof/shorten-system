package template

import (
	"fmt"
	"strings"
)

const emailVerificationCodeTemplate = `
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Email Verification</title>
</head>

<body style="margin:0;padding:0;background-color:#f4f6f8;font-family:Arial,sans-serif;">

<table width="100%" cellpadding="0" cellspacing="0" border="0" style="padding:40px 0;">
<tr>
<td align="center">

<table width="600" cellpadding="0" cellspacing="0" border="0"
style="background:#ffffff;border-radius:12px;padding:40px;">

<tr>
<td align="center">
<h1 style="margin:0;color:#111827;">
Verify Your Email Address
</h1>
</td>
</tr>

<tr>
<td style="padding-top:24px;color:#4b5563;font-size:16px;line-height:26px;">
Hello,

<br><br>

Thank you for using <strong>Shorten URL</strong>.

Please enter the verification code below to verify your email address.
</td>
</tr>

<tr>
<td align="center" style="padding:36px 0;">

<div style="
display:inline-block;
padding:18px 36px;
background:#f3f4f6;
border-radius:10px;
font-size:34px;
font-weight:700;
letter-spacing:10px;
color:#111827;
">
{{CODE}}
</div>

</td>
</tr>

<tr>
<td style="color:#4b5563;font-size:15px;line-height:24px;">
This verification code will expire in
<strong>5 minutes</strong>.
</td>
</tr>

<tr>
<td style="padding-top:12px;color:#4b5563;font-size:15px;line-height:24px;">
For security reasons, never share this code with anyone.
</td>
</tr>

<tr>
<td style="padding-top:12px;color:#4b5563;font-size:15px;line-height:24px;">
If you did not request this verification code, you can safely ignore this email.
</td>
</tr>

<tr>
<td style="padding-top:40px;">
<hr style="border:none;border-top:1px solid #e5e7eb;">
</td>
</tr>

<tr>
<td style="padding-top:16px;color:#9ca3af;font-size:13px;line-height:20px;">
Shorten URL Platform
<br>
This is an automated email. Please do not reply.
</td>
</tr>

</table>

</td>
</tr>
</table>

</body>
</html>
`

func EmailVerificationCode(
	code string,
) string {

	return strings.ReplaceAll(
		emailVerificationCodeTemplate,
		"{{CODE}}",
		fmt.Sprintf("%s", code),
	)
}
