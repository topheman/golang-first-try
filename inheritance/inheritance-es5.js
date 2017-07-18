function Animal(name) {
    this.name = name;
}
Animal.prototype.eat = function(foodName) {
    console.log(`${this.name} is eating ${foodName}`);
}

function Dog(name, isPet) {
    Animal.call(this, name);
    this.isPet = isPet;
}
Dog.prototype = Object.create(Animal.prototype);
Dog.prototype.constructor = Dog;
Dog.prototype.eat = function() {
    if (this.isPet) {
        return Animal.prototype.eat.call(this, 'dog food');
    }
    return console.log(`${this.name} is happy to eat!`);
}
Dog.prototype.bark = function() {
    console.log(`${this.name} barks!`);
}

const myDog = new Dog('Georges', true);
myDog.eat();
myDog.bark();

const regularDog = new Dog('Jules', false);
regularDog.eat();
regularDog.bark();