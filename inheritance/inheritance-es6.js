class Animal {
    constructor(name) {
        this.name = name;
    }
    eat(foodName) {
        console.log(`${this.name} is eating ${foodName}`);
    }
}

class Dog extends Animal {
    constructor(name, isPet) {
        super(name);
        this.isPet = isPet;
    }
    eat() {
        if (this.isPet) {
            return super.eat('dog food');
        }
        return console.log(`${this.name} is happy to eat!`);
    }
    bark() {
        console.log(`${this.name} barks!`);
    }
}

const myDog = new Dog('Georges', true);
myDog.eat();
myDog.bark();

const regularDog = new Dog('Jules', false);
regularDog.eat();
regularDog.bark();