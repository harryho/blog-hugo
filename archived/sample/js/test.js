// var a = "abc"
// console.log(a);

// this.a = "123";
// console.log(a);

// console.log(this);

// foo = 'abc';
// alert(foo); // abc

// this.foo = 'def';
// alert(foo); //

// console.log(  " 1==true : ", 1==true)
// console.log(  " ''==true : ", ''==true)
// console.log(  " '1'==true : ", '1'==true)
// console.log(  " \"\"==true : ", "1"==true)
// console.log(  " {}==true : ", [{}]==true)
// console.log(  " []==true : ", ['1']==true)

// console.log(  " 0==false : ", 0==false)
// console.log(  " -1==false : ", -1==false)
// console.log(  " ''==false : ", ''==false)
// console.log(  " '1'==false : ", '1'==false)
// console.log(  " \"\"==false : ", ""==false)
// console.log(  " {}==false : ", {}==false)
// console.log(  " []==false : ", []==false)

// console.log(  " null==false : ", null==false)
// console.log(  " undefined==false : ", undefined==false)
// console.log(  " undefined==null : ", null==undefined)

// var obj = {
//   _p: '_f',
//   _f: function() {
//     console.log(' f -> ', this);
//     // console.log(this.__proto__.constructor.name);
//     _ff = {
//       //   console.log(' ff ', this);
//       _p: '_ff',
//       _fff: function() {
//         this.p = '123';
//         console.log(' fff -> ', this);
//         // console.log(this.__proto__.constructor.name);
//       }
//     };
//     _ff._fff();
//   }
// };

// obj._f();

// console.log(' -------------------------------------- ');

// var obj2 = {
//   _p: 123,
//   _f: () => {
//     console.log(' f -> ', this);
//     // console.log(this.__proto__.constructor.name);
//     _ff = {
//       _p: '_ff',
//       _fff: () => {
//         this.p = '123';
//         console.log(' fff -> ', this);

//         //   console.log(this.__proto__.constructor.name);
//       }
//     };
//     _ff._fff();
//   }
// };

// obj2._f();

// ////////////////////////////////////////

// var obj3 = {
//   p: 'obj3',
//   toBeCalled: function() {
//     console.log(' this is toBeCalled ', this.p);
//   },
//   toBind: function(obj) {
//     obj.toBeCalled();
//     // console.log(' this is toBind ', obj);
//   }
// };

// var testBind = obj3.toBind;
// testBind(obj3);

// var obj4 = {
//   p: 'obj4',
//   toBeCalled: () => {
//     console.log(' this is toBeCalled ', this.p);
//   },
//   toBind: obj => {
//     obj.toBeCalled();
//     // console.log(' this is toBind ', obj);
//   }
// };

// var testBind2 = obj4.toBind;
// testBind2(obj4);

// var obj = {};

// const arr1= [1,2.20,-3.5]

// const arr2 =  [1.0,-3.5,2.2]

// console.log(  " arr1 = arr2 ? ", arr1.sort().toString()===arr2.sort().toString())
// console.log( obj1)
// console.log( obj2)

// function testVar(){
//     console.log(a)
//     console.log(b)
//     var b = 2
//     c = 123 // throws errors
// }

// var a = 1

// testVar()

// console.log(c)

// d = 13
// var e = 31
// console.log(this.e, this.d)
// console.log(e, d)
// delete this.d
// delete this.e
// console.log(this.e, this.d)
// console.log(e, d)

//  i = i + 1;

// console.log( i )
// console.log( this.i )

// var var1;
// let letVar;
// const constVar; // missing initialization

// function testVar() {
//     console.log( var1);
//     console.log( constVar);
//     console.log( letVar);
// }
// testVar()

// var v1 = "";
// var v1 = 123;

// let let1 = "";
// let let1 = 123;

// const c1 = "";
// c1 = 123;

// for ( var i = 0 ; i < 5 ; i++ ){
//     var x = 20;
//     console.log(i);
// }
// console.log( i );
// console.log( x );

// for ( ; i < 10 ; i++ ){
//     var i
//     console.log(i);
// }

// for ( let t = 0 ; t < 5 ; t++ ){
//     console.log( t);
//     let s = 100
// }
// console.log(t)
// console.log(s)

// const fs = require('fs');
// const csv = require('fast-csv');

// fs.createReadStream('test.csv')
//     .pipe(csv.parse({headers:true}))
//     .on('error', error => console.error(error))
//     .on('data', row => console.log(`ROW=${JSON.stringify(row)}`))
//     .on('end', rowCount => console.log(`Parsed ${rowCount} rows`));

// const csv = require('fast-csv');
// const fs = require('fs');
// const fileStream = fs.createReadStream('test.csv');
// const parser = csv.parse({ ignoreEmpty: true, headers: true, trim: true });

// fileStream
//   .pipe(parser)
//   // .validate(data => data.last_name && data.mobile.startsWith('04'))
//   .validate((row, cb) => {
//     setImmediate(() => cb(null, row.mobile.startsWith('04')));
//   })
//   .on('error', error => console.error(error))
//   .on('data', row => console.log(`Valid [row=${JSON.stringify(row)}]`))
//   .on('data-invalid', (row, rowNumber) => console.log(`Invalid [rowNumber=${rowNumber}] [row=${JSON.stringify(row)}]`))
//   .on('end', rowCount => console.log(`Parsed ${rowCount} rows`));

const { EOL } = require('os');
const csv = require('../../');

const CSV_STRING = ['firstName,lastName', 'bob,yukon', 'sally,yukon', 'timmy,yukon'].join(EOL);

const stream = csv
  .parse({ headers: true })
  .transform(data => ({
    firstName: data.firstName.toUpperCase(),
    lastName: data.lastName.toUpperCase(),
    properName: `${data.firstName} ${data.lastName}`
  }))
  .on('error', error => console.error(error))
  .on('data', row => console.log(JSON.stringify(row)))
  .on('end', rowCount => console.log(`Parsed ${rowCount} rows`));

stream.write(CSV_STRING);
stream.end();
