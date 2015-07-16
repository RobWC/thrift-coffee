namespace go basicservice
namespace py basicservice
namespace rb basicservice
namespace js basicservice

typedef i64 bigTime
//KeyValue service

struct Coffee {
    1: required string name;
    2: list<Creamer> creamers;
    3: list<Sweetner> sweetners;
    4: required Size size;
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
}

service CoffeeOrder {
    /*
     * Ping to check if server is alive
     */
    void ping(),
    void setValue(1:i32 key, 2:string value),
    i32 add(1:i32 a, 2:i32 b),
}
