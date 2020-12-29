const form = document.querySelector('.watson-form');
const operationSelect = document.querySelector('.operation-select');
const datatypeSelect = document.querySelector('.datatype-select');

const responseContainer = document.querySelector('.response-container');

const createError = err => {
  const div = document.createElement('div');
  div.classList.add('error');

  const strong = document.createElement('strong');
  strong.innerText = 'Error: ';

  const span = document.createElement('span');
  span.innerText = err;

  div.appendChild(strong);
  div.appendChild(span);

  responseContainer.innerHTML = div.outerHTML;
};

const createSuccess = data => {
  const pre = document.createElement('pre');
  const code = document.createElement('code');

  code.innerText = data;

  pre.appendChild(code);

  responseContainer.innerHTML = pre.outerHTML;
};

const toggleDatatype = () => {
  if (operationSelect.value === 'decode') {
    datatypeSelect.style.display = 'none';
  } else {
    datatypeSelect.style.display = 'block';
  }
};

const send = e => {
  e.preventDefault();

  responseContainer.innerHTML = 'Loading...';

  const operation = e.target.operation.value;
  const datatype = e.target.datatype.value;
  const input = e.target.input.value;

  const url = operation === 'encode'
    ? `/api/encode?type=${encodeURIComponent(datatype)}`
    : '/api/decode';

  return fetch(url, {
    method: 'POST',
    body: input,
    headers: {
      'Content-Type': 'application/json',
      'Accept': 'application/json'
    },
    credentials: 'same-origin'
  })
    .then(res => res.json())
    .then(json => {
      if (!json.ok) {
        return createError(json.error);
      }

      return createSuccess(json.data);
    })
    .catch(err => {
      console.error(err);
      return createError('Internal Server Error.');
    });
};

operationSelect.onchange = toggleDatatype;
form.onsubmit = send;
