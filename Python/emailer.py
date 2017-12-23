import smtplib
import sys
from email.MIMEMultipart import MIMEMultipart
from email.MIMEText import MIMEText
from email.MIMEBase import MIMEBase
from email import encoders
try:
    server = smtplib.SMTP('smtp.gmail.com', 587)
    username = 'your username'
    password = 'your password'

    server.ehlo()
    server.starttls()
    server.login(username,password)

    fromaddr = "your email address"
    toaddr = "your recipient address"

    msg = MIMEMultipart()
    msg['From'] = fromaddr
    msg['To'] = toaddr
    msg['Subject'] = 'your subject line'
    body = "the body of your email"
    msg.attach(MIMEText(body, 'plain'))
    # attaching a midi file to the email
    part = MIMEBase('application', 'octet-stream')
    part.set_payload(open("song.mid").read())
    encoders.encode_base64(part)
    part.add_header('Content-Disposition', 'attachment; filename= themidifile.mid')
    msg.attach(part)

    text = msg.as_string()
    server.sendmail(fromaddr, toaddr, text)
    server.close()
except:
    print "Email failed"
