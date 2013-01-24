K.<t> = QQ[]

# E.discriminant().factor() ...

def valuation(f,v,d):
    if v==1/t:
        return (f.degree()/d).ceil()*d-f.degree()
    assert(v.is_irreducible())
    ret = [n for (g,n) in f.factor() if v==g]
    if len(ret)==1:
        return ret[0]
    elif len(ret)==0:
        return 0
    assert(False)

def local_data(E,v):
    E = E.short_weierstrass_model()

    a = valuation(E.a4(),v,4)
    b = valuation(E.a6(),v,6)
    d = valuation(E.discriminant(),v,12)
    assert(a>=0 and b>=0 and d>=0)

    if d==0:
        return KodairaSymbol("I0")
    if a==0 and b==0:
        return KodairaSymbol("I%s"%d)
    if (a==2 and b>=3) or (a>=2 and b==3):
        assert(d>=6)
        assert(d==6 or (a==2 and b==3))
        return KodairaSymbol("I%s*"%(d-6)) 
    if d==2:
        assert(a>=1 and b==1)
        return KodairaSymbol("II")
    if d==3:
        assert(a==1 and b>=2)
        return KodairaSymbol("III")
    if d==4:
        assert(a>=2 and b==2)
        return KodairaSymbol("IV")
    if d==8:
        assert(a>=3 and b==4)
        return KodairaSymbol("IV*")
    if d==9:
        assert(a==3 and b>=5)
        return KodairaSymbol("III*")
    if d==10:
        assert(a>=4 and b==5)
        return KodairaSymbol("II*")
    assert(False)

def local_information(E):
    return [(v,local_data(E,v)) for (v,_) in list(E.discriminant().factor())+[(1/t,1)]]

E = EllipticCurve(K,[0,t^2+4*t+3,0,2*t+3,1])
P = [t-1,t,1/t]
print E; print [(p,local_data(E,p)) for p in P]
E = EllipticCurve(K,[-3*(t^4-12*t^3+14*t^2+12*t+1), 2*(t^6-18*t^5+75*t^4+75*t^2+18*t+1)])
print E; print [(p,local_data(E,p)) for p in P]
E = EllipticCurve(K,[-3*(t^2-3), t*(2*t^2-9)])
print E; print [(p,local_data(E,p)) for p in P]
print local_information(E)
