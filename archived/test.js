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

//////////////////////////////////////////

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



