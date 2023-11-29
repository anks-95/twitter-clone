interface FormValues {
  name:string;
  email: string;
  password: string;
}

interface ValidationErrors {
  name:string;
  email: string;
  password: string;
}


function Validation(values: FormValues): ValidationErrors {
  let errors: ValidationErrors = {
    name:'',
    email: '',
    password: '',
  };
  const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  const passwordPattern = /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])[a-zA-Z0-9]{8,}$/;


   if (values.name === "") {
    errors.name = "Email should not be empty";
  }{
    errors.name=" "
  }

  if (values.email === "") {
    errors.email = "Email should not be empty";
  } else if (!emailPattern.test(values.email)) {
    errors.email = "Email format is incorrect";
  }else{
    errors.email=" "
  }


  if (values.password === "") {
    errors.password = "Password should not be empty";
  } else if (!passwordPattern.test(values.password)) {
    errors.password = "Password format is incorrect";
  }else{
    errors.password=" "
  }


  return errors;
}


export default Validation;



// interface FormValues {
//   email: string;
//   password: string;
//   name: string;
// }

// interface ValidationErrors {
//   email: string;
//   password: string;
//   name: string;
// }


// function Validation(values: FormValues): ValidationErrors {
//   let errors: ValidationErrors = {
//     email: '',
//     password: '',
//     name:''
//   };
//   const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
//   const passwordPattern = /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])[a-zA-Z0-9]{8,}$/;

//   if (values.name === "") {
//     errors.name = "Name should not be empty";
//   } else{
//     errors.name=""
//   }


//   if (values.email === "") {
//     errors.email = "Email should not be empty";
//   } else if (!emailPattern.test(values.email)) {
//     errors.email = "Email format is incorrect";
//   }else{
//     errors.email=""
//   }

//   if (values.password === "") {
//     errors.password = "Password should not be empty";
//   } else if (!passwordPattern.test(values.password)) {
//     errors.password = "Password format is incorrect";
//   }else{
//     errors.password=""
//   }

//   return errors;
// }


// export default Validation;