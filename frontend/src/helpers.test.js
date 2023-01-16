import { NameToDateInt,NameFromPath,DateIntToName,GetDateRange } from "./helpers";

const arr = [{"Name":"A22110215424010.jpg","Date":"2022-11-02 19:42:44.396826519 +0000 UTC","Path":"/media/db/20221102/images/A22110215424010.jpg"},{"Name":"A22110219382210.jpg","Date":"2022-11-02 23:38:26.25391058 +0000 UTC","Path":"/media/db/20221102/images/A22110219382210.jpg"},{"Name":"A22110219392310.jpg","Date":"2022-11-02 23:39:27.140095115 +0000 UTC","Path":"/media/db/20221102/images/A22110219392310.jpg"},{"Name":"A22110219400510.jpg","Date":"2022-11-02 23:40:08.229569439 +0000 UTC","Path":"/media/db/20221102/images/A22110219400510.jpg"}]


test("GetDateRange", () => {
    expect(GetDateRange(arr)).toEqual(["A22110215424010.jpg","A22110219400510.jpg"]);  
})

test("DateIntToName", () => {
    expect(DateIntToName(22110212545710)).toBe("A22110212545710.jpg");
    expect(DateIntToName(22110213142710)).toBe("A22110213142710.jpg");
});

test("NameToDateInt", () => {
    expect(NameToDateInt("A22110212545710.jpg")).toBe(22110212545710);
    expect(NameToDateInt("A22110213142710.jpg")).toBe(22110213142710);

    // with Path
    expect(NameToDateInt("/media/db/20221102/images/A22110213142710.jpg")).toBe(22110213142710);
    expect(NameToDateInt("/media/db/20221102/images/A22110212545710.jpg")).toBe(22110212545710)
});


test("NameFromPath", () => {
    // with Path
    expect(NameFromPath("/media/db/20221102/images/A22110213142710.jpg")).toBe("A22110213142710.jpg");
    expect(NameFromPath("/media/db/20221102/images/A22110212545710.jpg")).toBe("A22110212545710.jpg")
});


// test("LargestArrayItem0", () => {
//     // with Path
//     expect(LargestArrayItem0([1,4])).toBe(false);
//     expect(LargestArrayItem0([99,3])).toBe(true)
//     expect(LargestArrayItem0([9,93])).toBe(false)
// });
