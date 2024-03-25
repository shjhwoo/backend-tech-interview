## 탬플릿 메서드

템플릿 메서드는 기초 클래스에서 알고리즘의 골격을 정의할 수 있도록 하는 행동 디자인 패턴입니다. 또 이 패턴은 자식 클래스들이 전체 알고리즘의 구조를 변경하지 않고도 기본 알고리즘의 단계들을 오버라이드할 수 있도록 합니다.

템플릿 메서드 패턴은 알고리즘을 일련의 단계들로 나누고, 이러한 단계들을 메서드들로 변환한 뒤, 단일 템플릿 메서드 내부에 이러한 메서드들에 대한 일련의 호출들을 넣으라고 제안합니다. 이러한 단계들은 abstract​(추상)​이거나 일부 디폴트​(기본값) 구현을 가질 것입니다. 알고리즘을 사용하기 위해 클라이언트는 자신의 자식 클래스를 제공해야 하고, 모든 추상 단계를 구현해야 하며, 필요하다면 (템플릿 메서드를 제외한) 선택적 단계 중 일부를 오버라이드​(재정의)​해야 합니다.

보시다시피 두 가지 유형의 단계들이 있습니다:

모든 자식 클래스는 추상 단계들을 구현해야 합니다.
선택적 단계들에는 이미 어떤 디폴트​(기본값) 구현이 있지만, 필요한 경우 이를 무시하고 오버라이드​(재정의) 할 수 있습니다.
훅이라는 또 다른 유형의 단계가 있습니다. 훅은 몸체가 비어 있는 선택적 단계입니다. 템플릿 메서드는 훅이 오버라이드 되지 않아도 작동합니다. 일반적으로 훅들은 알고리즘의 중요한 단계들의 전 또는 후에 배치되어 자식 클래스들에 알고리즘에 대한 추가 확장 지점들을 제공합니다.

```
// 추상 클래스는 템플릿 메서드를 정의합니다. 이 메서드는 일반적으로 원시 작업을
// 추상화하기 위해 호출로 구성된 어떤 알고리즘의 골격을 포함합니다. 구상 자식
// 클래스들은 이러한 작업을 구현하지만 템플릿 메서드 자체는 그대로 둡니다.
class GameAI is
    // 템플릿 메서드는 알고리즘의 골격을 정의합니다.
    method turn() is
        collectResources()
        buildStructures()
        buildUnits()
        attack()

    // 일부 단계들은 기초 클래스에서 바로 구현될 수 있습니다.
    method collectResources() is
        foreach (s in this.builtStructures) do
            s.collect()

    // 그리고 그중 일부는 추상으로 정의될 수 있습니다.
    abstract method buildStructures()
    abstract method buildUnits()

    // 한 클래스에는 여러 템플릿 메서드가 있을 수 있습니다.
    method attack() is
        enemy = closestEnemy()
        if (enemy == null)
            sendScouts(map.center)
        else
            sendWarriors(enemy.position)

    abstract method sendScouts(position)
    abstract method sendWarriors(position)

// 구상 클래스들은 기초 클래스의 모든 추상 작업을 구현해야 합니다. 하지만 템플릿
// 메서드 자체를 오버라이드해서는 안 됩니다.
class OrcsAI extends GameAI is
    method buildStructures() is
        if (there are some resources) then
            // 농장들, 막사들, 그리고 요새들을 차례로 건설하세요.

    method buildUnits() is
        if (there are plenty of resources) then
            if (there are no scouts)
                // 잡역인을 생성한 후 정찰병 그룹에 추가하세요.
            else
                // 하급 병사를 생성한 후 전사 그룹에 추가하세요.

    // …

    method sendScouts(position) is
        if (scouts.length > 0) then
            // 정찰병들을 위치로 보내세요.

    method sendWarriors(position) is
        if (warriors.length > 5) then
            // 전사들을 위치로 보내세요.

// 자식 클래스들은 디폴트 구현을 가진 일부 작업을 오버라이드할 수 있습니다.
class MonstersAI extends GameAI is
    method collectResources() is
        // 몬스터들은 자원을 모으지 않습니다.

    method buildStructures() is
        // 몬스터들은 건물을 짓지 않습니다.

    method buildUnits() is
        // 몬스터들은 유닛들을 생성하지 않습니다.
```

### 장단점

- 장점
  클라이언트들이 대규모 알고리즘의 특정 부분만 오버라이드하도록 하여 그들이 알고리즘의 다른 부분에 발생하는 변경에 영향을 덜 받도록 할 수 있습니다.
  중복 코드를 부모 클래스로 가져올 수 있습니다.

* 단점
  일부 클라이언트들은 알고리즘의 제공된 골격에 의해 제한될 수 있습니다.
  당신은 자식 클래스를 통해 디폴트 단계 구현을 억제하여 리스코프 치환 원칙을 위반할 수 있습니다.
  템플릿 메서드들은 단계들이 더 많을수록 유지가 더 어려운 경향이 있습니다.
