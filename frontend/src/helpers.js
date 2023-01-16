import axios from "axios"

export const FetchPics = async (setPics) => {
    axios.get('http://192.168.1.209:30099/list')
        .then(res => {
            res.data.map(pic => pic.Active = false)
            setPics(res.data.reverse())
        })
};

export const SetActive = (name, pics, setPics) => {
    pics.map(pic => {
        if (pic.Name === name) {
            pic.Active = true
        }
        return true
    })
    setPics(pics)
}

export const SetUnActive = (name, pics, setPics) => {
    pics.map(pic => {
        if (pic.Name === name) {
            pic.Active = false;
        }
        return true
    })
    setPics(pics)
}

export const DeletePic = (e, path, pics, setPics, filterPics, date) => {
    if (e.detail === 2) {
        axios.get('http://192.168.1.209:30099/delete/file' + path)
            .then(res => {
                if (res.data) {
                    setPics(pics.filter(pic => pic.Path !== path));
                }
            })
    } else if (e.detail === 4) {
        axios.get(`http://192.168.1.209:30099/delete/cascade?begin=${GetDateRange(filterPics(pics, date))[0]}&end=${GetDateRange(filterPics(pics, date))[1]}`)
            .then(res => {
                if (res.data) {
                    setPics(pics.filter(pic => NameToDateInt(pic.Path) >= GetDateRange(filterPics(pics, date))[0] && NameToDateInt(pic.Path) <= GetDateRange(filterPics(pics, date))[1]));
                }
            })
    }
}

export const FilterPics = (pics, date) => {
    if (date !== "") {
        return pics.filter(pic => pic.Name.substring(0, date.length + 1) === "A" + date)
    } else {
        return pics
    }
}

export function NameToDateInt(name) {
    return parseInt(name.split("A")[1] + name.split(".")[0]);
}

export function DateIntToName(dateint) {
    return "A" + dateint + ".jpg";
}

export function NameFromPath(path) {
    return path.split("/").pop();
}

export function GetDateRange(arr) {
    if (Array.isArray(arr)) {
        let dateIntArray = arr.map(pic => NameToDateInt(pic.Name));
        console.log("DateIntArray ", JSON.stringify(dateIntArray, undefined, 2))
        dateIntArray.sort((a, b) => a - b);
        return [DateIntToName(dateIntArray[0]), DateIntToName(dateIntArray[dateIntArray.length - 1])];
    } else {
        return []
    }
}