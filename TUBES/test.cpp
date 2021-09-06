#include "test.h"

void createList(List_paper &L) {
    first(L) = NULL;
    last(L) = NULL;

}

address_child alokasi(infotype_child x) {
    address_paper P = new elmlist_paper;
    info(P) = x;
    next(P) = NULL;
    prev(P) = NULL;
    return P;
}

void insertFirst(List_paper &L, address_paper P) {
    if(first(L) == NULL) {
        last(L) = P;
        first(L) = P;
    } else {
        next(P) = first(L);
        prev(first(L)) = P;
        first(L) = P;
    }
}

void printInfo(List_paper L) {
    address_paper P = first(L);
    while(P !=NULL) {
        cout<<"->"<<info(P)<<endl;
        P = next(P);
    }
}


address_child findElm(List_paper L, infotype_paper x) {
    address_paper P = first(L);
    while(P != NULL) {
        if(info(P)==x) {
            return P;
        }
        P = next(P);
    }
    return NULL;
}

void insertAfter(address_paper &Prec, address_paper P) {
    prev(next(Prec)) = P;
    next(P) = next(Prec);
    prev(P) = Prec;
    next(Prec) = P;
}
