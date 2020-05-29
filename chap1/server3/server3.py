import io
import sys
from flask import Flask, request, render_template

app = Flask(__name__)


@app.route('/')
def hello_world():
    out = io.StringIO()
    print(f'{request.method} {request.url} {request.scheme}', file=out)
    for k, v in sorted(request.headers):
        print(f'Header[{k!r}] = [{v!r}]', file=out)
    print(f'Host = {request.host!r}', file=out)
    print(f'RemoteAddr = {request.remote_addr!r}', file=out)
    return render_template('server3.html', text=out.getvalue())


if __name__ == "__main__":
    app.run(debug=True)
