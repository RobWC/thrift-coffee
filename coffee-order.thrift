namespace go CoffeeService
namespace py CoffeeService
namespace rb CoffeeService
namespace js CoffeeService

struct Coffee {
    1: required string name;
    2: list<Creamer> creamers;
    3: list<Sweetner> sweetners;
    4: Size size;
    5: bool iced;
    6: i32 temperature;
}

struct Size {
    1: required string name;
    2: required i32 ounces;
}

struct Creamer {
    1: required string name;
    2: string flavor;
}

struct Sweetner {
    1: required string name;
    2: required i32 amount;
    3: required string units;
    4: optional string flavor;
}

exception InvalidCoffee {
    1: required i32 orderID,
    2: required string error
}

service CoffeeOrder {
    /*
     * Ping to check if server is alive
     */
    void ping(),
    /*
     * orderCoffee
     */
    i32 orderCoffee(1:Coffee coffee) throws (1:InvalidCoffee orderError),
}
